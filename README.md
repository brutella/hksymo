# hksymo

This project is an implementation of a HomeKit bridge for the [Fronius Symo](http://www.fronius.com/cps/rde/xchg/fronius_international/hs.xsl/83_28694_ENG_HTML.htm) using [HomeControl](https://github.com/brutella/hc) and [fronius](https://github.com/brutella/fronius).

# Installation

## Build

Build `hksymod.go` using `go build hksymod.go` or use the Makefile to build for Raspberry Pi

    make rpi

## Run

Simply execute the daemon with `./hksymod`. By default the daemon tries to connect to the Fronius Symo by using the predefined IP address `169.254.0.180`. 

If the inverter has a fixed or dynamic IP address, you should use that address

    ./hksymod --host=10.0.0.5


## HomeKit Client

You need an iOS app to control HomeKit accessories. 
You can use [Home][home] which runs on iPhone, iPad and Apple Watch.

[home]: http://selfcoded.com/home
[home-getting-started]: http://selfcoded.com/home/getting-started

## TODO

- Show all values provided by the API

# Contact

Matthias Hochgatterer

Github: [https://github.com/brutella](https://github.com/brutella)

Twitter: [https://twitter.com/brutella](https://twitter.com/brutella)

# License

hksymod is available under a non-commercial license. See the LICENSE file for more info.