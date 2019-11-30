UNAME := $(shell uname)

all: build

build: 
	go build -ldflags="-s -w" -o bin/dive_core ./dive.go
	upx bin/dive_core

install:
	make install_$(UNAME)

install_Darwin:
	rm -f /usr/local/bin/dive
	rm -f /usr/local/bin/dive_core
	cp -rf bin/dive /usr/local/bin/dive
	cp -rf bin/dive_core /usr/local/bin/dive_core
	chmod +x /usr/local/bin/dive

install_Linux:
	cp ./src/dive.py /usr/bin/dive
	chmod +x /usr/bin/dive
