package main

import (

	//"flag"
	//"time"
	//"encoding/json"

	"time"

	"github.com/asticode/go-astilectron"
	//"github.com/asticode/go-astilectron-bootstrap"
	"github.com/asticode/go-astilog"
	//"github.com/pkg/errors"
)

func main() {

	// Initialize astilectron
	var a, _ = astilectron.New(astilectron.Options{
		AppName:            "<your app name>",
		AppIconDefaultPath: "<your .png icon>",  // If path is relative, it must be relative to the data directory
		AppIconDarwinPath:  "<your .icns icon>", // Same here
		BaseDirectoryPath:  "<where you want the provisioner to install the dependencies>",
		VersionAstilectron: "<version of Astilectron to utilize such as `0.33.0`>",
		VersionElectron:    "<version of Electron to utilize such as `4.0.1` | `6.1.2`>",
	})
	defer a.Close()

	// Start astilectron
	a.Start()

	// Blocking pattern
	a.Wait()
	// Create a new window
	var w, _ = a.NewWindow("http://127.0.0.1:4000", &astilectron.WindowOptions{
		Center: astilectron.PtrBool(true),
		Height: astilectron.PtrInt(600),
		Width:  astilectron.PtrInt(600),
	})
	w.Create()
	// Open dev tools
	w.OpenDevTools()
	// Add a listener on Astilectron
	a.On(astilectron.EventNameAppCrash, func(e astilectron.Event) (deleteListener bool) {
		astilog.Error("App has crashed")
		return
	})

	// Add a listener on the window
	w.On(astilectron.EventNameWindowEventResize, func(e astilectron.Event) (deleteListener bool) {
		astilog.Info("Window resized")
		return
	})

	// Play with the window
	w.Resize(200, 200)
	time.Sleep(time.Second)
	w.Maximize()

}
