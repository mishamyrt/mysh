UNAME := $(shell uname)
LD_FLAGS=-s -w
BUILD_CMD=go build -ldflags="$(LD_FLAGS)"

all: build

build_all: build_Darwin build_Linux

build:
	mkdir -p dist
	mkdir -p dist/macOS
	mkdir -p dist/linux64
	make build_$(UNAME)

build_Darwin:
	env GOOS=darwin GOARCH=amd64 $(BUILD_CMD) -o dist/macOS/dive_core ./dive.go
	upx dist/macOS/dive_core

build_Linux:
	env GOOS=linux GOARCH=amd64 $(BUILD_CMD) -o dist/linux64/dive_core ./dive.go
	upx dist/linux64/dive_core

install:
	make install_$(UNAME)

install_Darwin:
	make uninstall_unix
	cp -rf bin/dive_unix /usr/local/bin/dive
	cp -rf dist/macOS/dive_core /usr/local/bin/dive_core
	make chmod_unix

chmod_unix:
	chmod +x /usr/local/bin/dive_core
	chmod +x /usr/local/bin/dive

uninstall_unix:
	rm -f /usr/local/bin/dive
	rm -f /usr/local/bin/dive_core

install_Linux:
	make uninstall_unix
	cp -rf bin/dive_unix /usr/local/bin/dive
	cp -rf dist/linux64/dive_core /usr/local/bin/dive_core
	make chmod_unix
