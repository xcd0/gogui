
addpath=`./gotron-builder/addpath.sh`

ifeq ($(shell uname -o),Msys)
os=win
else ifeq ($(shell uname),Darwin)
os=mac
else ifeq ($(shell uname),Linux)
os=linux
endif
dirname=$(os)-unpacked
binname=`ls -F ./dist/$(dirname) | grep -v / | sed 's/\*//g'`
bin=./dist/$(dirname)/$(binname)

build:
	go build
	$(addpath)
	gotron-builder

run:
ifeq ($(shell uname -o),Msys)
	make run-win
else ifeq ($(shell uname),Darwin)
	make run-mac
else ifeq ($(shell uname),Linux)
	make run-linux
endif

run-win:
	$(bin)
run-mac:
	open $(bin)
run-linux:
	$(bin)

