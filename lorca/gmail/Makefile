
BIN=gmail

all: build

build:
	rsrc -ico $(BIN).ico -o $(BIN).syso
ifeq ($(shell uname -o),Msys)
	go build -o $(BIN).exe -ldflags="-H windowsgui"
else
	go build -o $(BIN)
endif

run:
	./$(BIN)

