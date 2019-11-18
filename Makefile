UNAME := $(shell uname)

install:
	make install_$(UNAME)

install_Darwin:
	cp ./src/dive.py /usr/local/bin/dive
	chmod +x /usr/local/bin/dive

install_Linux:
	cp ./src/dive.py /usr/bin/dive
	chmod +x /usr/bin/dive
