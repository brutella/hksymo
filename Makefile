GO ?= go
FILE = daemon/hksymod.go
bbb:
	GOOS=linux GOARCH=arm GOARM=7 $(GO) build $(FILE)

rpi:
	GOOS=linux GOARCH=arm GOARM=6 $(GO) build $(FILE)