GO ?= go
FILE = daemon/hksymod.go

rpi:
	GOOS=linux GOARCH=arm GOARM=6 $(GO) build $(FILE)