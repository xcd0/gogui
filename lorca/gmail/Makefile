
all: build

build:
	rsrc -ico gmail.ico -o gmail.syso
ifeq ($(shell uname -o),Msys)
	go build -o gmail.exe -ldflags="-H windowsgui"
else
	go build -o gmail
endif

run:
	./gmail

clean:
	rm -rf gmail.exe gmail.syso
