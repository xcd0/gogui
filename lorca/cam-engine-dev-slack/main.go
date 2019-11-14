package main

import (
	"github.com/zserge/lorca"
	"log"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	dir := user.HomeDir + "/.config"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, 0777)
	}
	dir = user.HomeDir + "/.config/lorca"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, 0777)
	}
	dir = user.HomeDir + "/.config/lorca/cam-engine-dev-slack"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, 0777)
	}

	os.MkdirAll(dir, os.ModePerm)
	ui, _ := lorca.New("https://cam-engine-dev.slack.com", dir, 1000, 1000)
	defer ui.Close()
	<-ui.Done()
}
