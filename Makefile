VERSION := $(shell cat CHANGELOG.md | grep -m 1 "\#\#"  | cut -d' ' -f2 | cut -d "[" -f2 | cut -d "]" -f1)
UNAME := $(shell uname)
DATE := $(shell date)
GIT_COMMIT := $(shell git rev-list -1 HEAD)
CMD_PACKAGE = github.com/mishamyrt/mysh/v1/cmd
LD_FLAGS=-s -w -X "$(CMD_PACKAGE).GitCommit=$(GIT_COMMIT)" -X "$(CMD_PACKAGE).Version=$(VERSION)" -X "$(CMD_PACKAGE).BuildTime=$(DATE)"
BUILD_CMD=go build -ldflags='$(LD_FLAGS)'

DEB_VERSION := $(shell echo "$(VERSION)" | sed 's/\(.*\)\./\1-/')
DEB_FOLDER=dist/mysh_$(DEB_VERSION)-amd64
DARWIN_FOLDER=dist/mysh_$(VERSION)_darwin_amd64

DARWIN_BINARY=$(DARWIN_FOLDER)/core_mysh
LINUX_BINARY=$(DEB_FOLDER)/usr/local/bin/core_mysh

all: brew_package deb_package

lint:
	golangci-lint run -E lll -E misspell -E prealloc -E stylecheck -E gocritic

brew_package: build_darwin
	mkdir -p $(DARWIN_FOLDER)
	cp LICENSE $(DARWIN_FOLDER)/
	cp entrypoint/mysh_unix $(DARWIN_FOLDER)/mysh
	cp completions/mysh.bash $(DARWIN_FOLDER)/
	cp completions/mysh.fish $(DARWIN_FOLDER)/
	chmod +x $(DARWIN_FOLDER)/*
	tar czf $(DARWIN_FOLDER).tar.gz --directory=$(DARWIN_FOLDER)/ .
	rm -rf $(DARWIN_FOLDER)

deb_package: build_linux
	mkdir $(DEB_FOLDER)/usr/share/
	mkdir $(DEB_FOLDER)/usr/share/bash-completion
	mkdir $(DEB_FOLDER)/usr/share/bash-completion/completions
	mkdir $(DEB_FOLDER)/usr/share/fish
	mkdir $(DEB_FOLDER)/usr/share/fish/vendor_completions.d
	mkdir $(DEB_FOLDER)/DEBIAN
	cp entrypoint/mysh_unix $(DEB_FOLDER)/usr/local/bin/mysh
	sed  's/__VERSION__/$(DEB_VERSION)/g' debian/control.template > $(DEB_FOLDER)/DEBIAN/control
	cp completions/mysh.bash $(DEB_FOLDER)/usr/share/bash-completion/completions/mysh
	cp completions/mysh.fish $(DEB_FOLDER)/usr/share/fish/vendor_completions.d/mysh
	chmod +x $(DEB_FOLDER)/usr/local/bin/*
	chmod +x $(DEB_FOLDER)/usr/share/bash-completion/completions/mysh
	chmod +x $(DEB_FOLDER)/usr/share/fish/vendor_completions.d/mysh
	dpkg-deb --build $(DEB_FOLDER)
	rm -rf $(DEB_FOLDER)

build_darwin:
	env GOOS=darwin GOARCH=amd64 $(BUILD_CMD) -o $(DARWIN_BINARY) ./mysh.go
	upx $(DARWIN_BINARY)

build_linux:
	env GOOS=linux GOARCH=amd64 $(BUILD_CMD) -o $(LINUX_BINARY) ./mysh.go
	upx $(LINUX_BINARY)

install:
	make install_$(UNAME)

install_Darwin:
	make uninstall_unix
	cp -rf entrypoint/mysh_unix /usr/local/bin/mysh
	cp -rf dist/macOS/mysh_core /usr/local/bin/mysh_core
	make chmod_unix

install_fish:
	mkdir -p ~/.config/fish/completions
	rm -f ~/.config/fish/completions/mysh.fish
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
	cp -rf entrypoint/mysh_unix /usr/local/bin/mysh
	cp -rf dist/linux64/mysh_core /usr/local/bin/mysh_core
	make chmod_unix
