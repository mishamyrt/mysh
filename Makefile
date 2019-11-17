UNAME := $(shell uname)

install:
	make install_$(UNAME)

install_Darwin:
	cp ./dive.py /usr/local/bin/dive
	chmod +x /usr/local/bin/dive

install_Linux:
	cp ./dive.py /usr/bin/dive
	chmod +x /usr/bin/dive
