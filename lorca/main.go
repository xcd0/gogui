package main

import (
	"github.com/zserge/lorca"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	dir := user.HomeDir + "/.config"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, 0777)
	}
	dir = user.HomeDir + "/.config/lorca"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, 0777)
	}
	dir = user.HomeDir + "/.config/lorca/app"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, 0777)
	}

	os.MkdirAll(dir, os.ModePerm)
	ui, _ := lorca.New("https://mail.google.com/", dir, 800, 1000)
	defer ui.Close()
	<-ui.Done()
}
