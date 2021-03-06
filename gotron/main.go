package main

import (
	"bytes"
	"fmt"
	"github.com/Equanox/gotron"
)

func main() {
	// browser window instanceを生成する
	window, err := gotron.New("app")
	if err != nil {
		panic(err)
	}

	// デフォルトのwindowサイズとタイトルを変更する
	window.WindowOptions.Width = 800
	window.WindowOptions.Height = 800
	window.WindowOptions.Title = "Gotron"

	// browser window を起動
	// window objectを使ってElectronBrowserWindow をコントロールするために
	// websocketを使って go と nodejsのブリッジを確立するよう
	done, err := window.Start()
	if err != nil {
		panic(err)
	}

	// dev toolsを使う場合は下記をコメントアウトする(せっかくなのでdevtools使ってみる)
	//window.OpenDevTools()

	window.On(
		&gotron.Event{Event: "event-name"},
		func(bin []byte) {
			b := []byte(bin)
			buf := bytes.NewBuffer(b)
			fmt.Println(buf)
		})

	// アプリケーションがcloseするのを待つ
	<-done
}
