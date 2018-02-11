package main

import (
	"flag"
	"github.com/brutella/fronius"
	"github.com/brutella/hc"
	"github.com/brutella/hc/accessory"
	"github.com/brutella/hksymo"

	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"time"
)

// from http://stackoverflow.com/a/16930649/424814
func TimeoutDialer(cTimeout time.Duration, rwTimeout time.Duration) func(net, addr string) (c net.Conn, err error) {
	return func(netw, addr string) (net.Conn, error) {
		conn, err := net.DialTimeout(netw, addr, cTimeout)
		if err != nil {
			return nil, err
		}
		conn.SetDeadline(time.Now().Add(rwTimeout))
		return conn, nil
	}
}

func NewTimeoutClient(connectTimeout time.Duration, readWriteTimeout time.Duration) *http.Client {

	return &http.Client{
		Transport: &http.Transport{
			Dial: TimeoutDialer(connectTimeout, readWriteTimeout),
		},
	}
}

func get() (inv fronius.InverterSystemResponse, err error) {
	client := NewTimeoutClient(5*time.Second, 5*time.Second)

	if simulate == true {
		s := fronius.NewSymoSimulator()
		defer s.Stop()
		url, _ := url.Parse(s.URL())
		resp, err := client.Get(fronius.SystemRealtimeDataRequestURL(url.Host))

		if err != nil {
			log.Fatal(err)
		}

		inv, err = fronius.NewInverterSystemResponse(resp)
	} else {
		var resp *http.Response
		resp, err = client.Get(fronius.SystemRealtimeDataRequestURL(host))

		if err == nil {
			inv, err = fronius.NewInverterSystemResponse(resp)
		}
	}

	return inv, err
}

func update() {
	inv, err := get()

	if err != nil {
		log.Println(err)

		if transport != nil {
			transport.Stop()
			transport = nil
		}
		return
	}

	pow := fronius.SystemCurrentPower(inv)
	today := fronius.SystemEnergyToday(inv)
	year := fronius.SystemEnergyThisYear(inv)
	total := fronius.SystemEnergyTotal(inv)
	acc.Inverter.Current.SetValue(int(pow.Value))
	acc.Inverter.Today.SetValue(int(today.Value))
	acc.Inverter.Year.SetValue(int(year.Value))
	acc.Inverter.Total.SetValue(int(total.Value))

	if transport == nil {
		var err error
		transport, err = hc.NewIPTransport(config, acc.Accessory)
		if err != nil {
			log.Println(err)
		}

		go func() {
			transport.Start()
		}()
	}
}

var transport hc.Transport
var acc *symo.Accessory
var host string
var simulate bool
var config hc.Config

func main() {
	var (
		hostArg     = flag.String("host", fronius.SymoHostClassA, "Host; default 169.254.0.180")
		refreshArg  = flag.Int("refresh", 10, "Refresh in seconds; default 10")
		simulateArg = flag.Bool("simulate", false, "Simulate Fronius symo; default true")
	)

	flag.Parse()
	host = *hostArg
	simulate = *simulateArg

	refresh := time.Duration(*refreshArg) * time.Second

	info := accessory.Info{
		Name:         "Symo",
		Manufacturer: "Fronius",
	}

	acc = symo.NewAccessory(info)

	var timer *time.Timer
	timer = time.AfterFunc(refresh, func() {
		update()
		timer.Reset(refresh)
	})

	hc.OnTermination(func() {
		if transport != nil {
			<-transport.Stop()
		}
		timer.Stop()
		os.Exit(1)
	})

	update()

	select {}
}
