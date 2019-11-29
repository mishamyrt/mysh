UNAME := $(shell uname)

all: build

build: 
	go build -ldflags="-s -w" -o bin/dive-core ./src/dive

install:
	rm /usr/local/share/dive/dive-core
	ln bin/dive-core /usr/local/share/dive/dive-core 

compress:
	upx --brute bin/dive-core

# install:
# 	make install_$(UNAME)

# install_Darwin:
# 	cp ./src/dive.py /usr/local/bin/dive
# 	chmod +x /usr/local/bin/dive

# install_Linux:
# 	cp ./src/dive.py /usr/bin/dive
# 	chmod +x /usr/bin/dive
