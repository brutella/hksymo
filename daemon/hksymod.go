package main

import (
	"flag"
	"fmt"
	"github.com/brutella/fronius"
	"github.com/brutella/hc/hap"
	"github.com/brutella/hc/model"
	"github.com/brutella/hksymo"
	"github.com/brutella/log"

	slog "log"
	"net"
	"net/http"
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
	if simulate == true {
		s := fronius.NewSymoSimulator()
		defer s.Stop()
		resp, err := fronius.GetSystemRealtimeData(s.URL())

		if err != nil {
			slog.Fatal(err)
		}

		inv, err = fronius.NewInverterSystemResponse(resp)
	} else {
		client := NewTimeoutClient(5*time.Second, 5*time.Second)
		var resp *http.Response
		resp, err = client.Get(fronius.SystemRealtimeDataRequestURL(fronius.SymoHostClassA))

		if err != nil {
			if transport != nil {
				transport.Stop()
				transport = nil
			}
		} else {
			inv, err = fronius.NewInverterSystemResponse(resp)
		}
	}

	return inv, err
}

func update() {
	log.Verbose = false

	inv, err := get()

	if err != nil {
		slog.Println(err)
		return
	}

	fmt.Printf("current power: %v\n", fronius.SystemCurrentPower(inv))
	fmt.Printf("today: %v\n", fronius.SystemEnergyToday(inv))
	fmt.Printf("this year: %v\n", fronius.SystemEnergyThisYear(inv))
	fmt.Printf("total: %v\n", fronius.SystemEnergyTotal(inv))

	pow := fronius.SystemCurrentPower(inv)
	today := fronius.SystemEnergyToday(inv)
	year := fronius.SystemEnergyThisYear(inv)
	total := fronius.SystemEnergyTotal(inv)
	generator.SetCurrentPower(pow.Value)
	generator.SetTodayEnergy(today.Value)
	generator.SetYearEnergy(year.Value)
	generator.SetTotalEnergy(total.Value)

	if transport == nil {
		var err error
		transport, err = hap.NewIPTransport("00102003", generator.Accessory)
		if err != nil {
			slog.Println(err)
		}

		go func() {
			transport.Start()
		}()
	}
}

var transport hap.Transport
var generator *symo.Generator
var host string
var simulate bool

func main() {
	var (
		hostArg     = flag.String("host", fronius.SymoHostClassA, "Host; default http://169.254.0.180")
		refreshArg  = flag.Int("refresh", 10, "Refresh in seconds; default 10")
		simulateArg = flag.Bool("simulate", true, "Simulate Fronius symo; default true")
	)

	flag.Parse()

	host = *hostArg
	simulate = *simulateArg
	refresh := time.Duration(*refreshArg) * time.Second

	info := model.Info{
		Name:         "Symo",
		Manufacturer: "Fronius",
	}

	generator = symo.NewGenerator(info)

	var timer *time.Timer
	timer = time.AfterFunc(refresh, func() {
		update()
		timer.Reset(refresh)
	})

	hap.OnTermination(func() {
		if transport != nil {
			transport.Stop()
		}
		timer.Stop()
		os.Exit(1)
	})

	update()

	select {}
}
