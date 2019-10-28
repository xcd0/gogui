package main

import "github.com/zserge/webview"

func main() {
	// Open wikipedia in a 800x600 resizable window
	webview.Open("Gmail",
		"https://mail.google.com", 800, 600, true)
}
