GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

VERSION=$(shell git describe --exact-match --tags 2>/dev/null)

BUILD_DIR=build

PACKAGE_UVRDUMP=uvrdump-$(VERSION)_linux_armhf
PACKAGE_UVRINFLUX=uvrinflux-$(VERSION)_linux_armhf

all: test build
build:
	$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) -i $(BUILD_SRC)

test:
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -rf $(BINARY_NAME)

package-uvrdump: build-uvrdump
	tar -cvzf $(PACKAGE_UVRDUMP).tar.gz -C $(BUILD_DIR) $(PACKAGE_UVRDUMP)

package-uvrinflux: build-uvrinflux
	tar -cvzf $(PACKAGE_UVRINFLUX).tar.gz -C $(BUILD_DIR) $(PACKAGE_UVRINFLUX)

build-uvrdump:
	GOOS=linux GOARCH=arm GOARM=6 $(GOBUILD) -o $(BUILD_DIR)/uvrdump-$(VERSION)_linux_armhf/usr/bin/uvrdump -i cmd/uvrdump/main.go

build-uvrinflux:
	GOOS=linux GOARCH=arm GOARM=6 $(GOBUILD) -o $(BUILD_DIR)/uvrinflux-$(VERSION)_linux_armhf/usr/bin/uvrinflux -i cmd/uvrinflux/main.go