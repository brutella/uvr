GO ?= go

rpi:
	GOOS=linux GOARCH=arm GOARM=7 $(GO) build cmd/uvrdump.go