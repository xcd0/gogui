
BIN=cam-engine-dev-slack
ICO=app-release

all: build

build:
	rsrc -ico $(ICO).ico -o $(ICO).syso
ifeq ($(shell uname -o),Msys)
	go build -o $(BIN).exe -ldflags="-H windowsgui"
else
	go build -o $(BIN)
endif

run:
	./$(BIN)

