UNAME := $(shell uname)
DATE := $(shell date)
GIT_COMMIT := $(shell git rev-list -1 HEAD)
VERSION=0.1.0
DEB_VERSION=0.1-0
DEB_FOLDER=mysh_$(DEB_VERSION)
LD_FLAGS=-s -w -X "main.GitCommit=$(GIT_COMMIT)" -X "main.Version=$(VERSION)" -X "main.BuildTime=$(DATE)"
BUILD_CMD=go build -ldflags='$(LD_FLAGS)'

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
	rm -f ~/.config/fish/completions/mysh.fish
	cp completions/mysh.fish ~/.config/fish/completions
	chmod +x ~/.config/fish/completions/mysh.fish

deb:
	mkdir $(DEB_FOLDER)
	mkdir $(DEB_FOLDER)/usr
	mkdir $(DEB_FOLDER)/usr/share/
	mkdir $(DEB_FOLDER)/usr/share/mysh
	mkdir $(DEB_FOLDER)/usr/share/bash-completion
	mkdir $(DEB_FOLDER)/usr/share/bash-completion/completions
	mkdir $(DEB_FOLDER)/usr/share/fish
	mkdir $(DEB_FOLDER)/usr/share/fish/vendor_completions.d
	mkdir $(DEB_FOLDER)/usr/local/
	mkdir $(DEB_FOLDER)/usr/local/bin
	mkdir $(DEB_FOLDER)/DEBIAN
	cp dist/linux64/mysh_core $(DEB_FOLDER)/usr/share/mysh
	cp cmd/mysh_linux $(DEB_FOLDER)/usr/local/bin/mysh
	cp debian/control $(DEB_FOLDER)/DEBIAN/control
	cp debian/postinst $(DEB_FOLDER)/DEBIAN/postinst
	cp completions/mysh.bash $(DEB_FOLDER)/usr/share/bash-completion/completions/mysh
	cp completions/mysh.fish $(DEB_FOLDER)/usr/share/fish/vendor_completions.d/mysh
	chmod +x $(DEB_FOLDER)/DEBIAN/postinst
	chmod +x $(DEB_FOLDER)/usr/local/bin/mysh
	chmod +x $(DEB_FOLDER)/usr/share/bash-completion/completions/mysh
	chmod +x $(DEB_FOLDER)/usr/share/fish/vendor_completions.d/mysh
	dpkg-deb --build $(DEB_FOLDER)

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
