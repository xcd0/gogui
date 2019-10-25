#!/bin/bash

script_dir=$(cd $(dirname $0); pwd)

if [ ! -e $script_dir/gotron-builder-* ]; then
	url="https://github.com/Equanox/gotron/releases/download/v0.2.23/gotron-builder-amd64-"
	if [ "$(uname)" == "Darwin" ]; then
		os=darwin
	elif [ "$(expr substr $(uname -s) 1 5)" == "Linux" ]; then
		os=linux
	elif [ "$(expr substr $(uname -s) 1 4)" == "MSYS" ]; then
		os=win.exe
	else
		echo "Your platform ($(uname -s)) is not supported."
		exit 1
	fi
	wget -q ${url}${os}
	ln -sf $script_dir/gotron-builder-amd64-${os} $script_dir/gotron-builder
	echo "export PATH=\$PATH:$script_dir" > ~/.bashrc
fi

echo "export PATH=\$PATH:$script_dir"

