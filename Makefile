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
	env GOOS=darwin GOARCH=amd64 $(BUILD_CMD) -o dist/macOS/mysh_core ./mysh.go
	upx dist/macOS/mysh_core

build_Linux:
	env GOOS=linux GOARCH=amd64 $(BUILD_CMD) -o dist/linux64/mysh_core ./mysh.go
	upx dist/linux64/mysh_core

install:
	make install_$(UNAME)

install_Darwin:
	make uninstall_unix
	cp -rf cmd/mysh_unix /usr/local/bin/mysh
	cp -rf dist/macOS/mysh_core /usr/local/bin/mysh_core
	make chmod_unix

install_fish:
	mkdir -p ~/.config/fish/completions
	cp completions/mysh.fish ~/.config/fish/completions
	chmod +x ~/.config/fish/completions/mysh.fish


chmod_unix:
	chmod +x /usr/local/bin/mysh_core
	chmod +x /usr/local/bin/mysh

uninstall_unix:
	rm -f /usr/local/bin/mysh
	rm -f /usr/local/bin/mysh_core

install_Linux:
	make uninstall_unix
	cp -rf cmd/mysh_unix /usr/local/bin/mysh
	cp -rf dist/linux64/mysh_core /usr/local/bin/mysh_core
	make chmod_unix
