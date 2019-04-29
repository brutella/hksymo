GO ?= go
FILE = daemon/hksymod.go

export GO111MODULE=on

rpi:
	GOOS=linux GOARCH=arm GOARM=6 $(GO) build $(FILE)