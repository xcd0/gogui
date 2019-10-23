package main

import (
	"github.com/Equanox/gotron"
)

func main() {
	// browser window instanceを生成する
	window, err := gotron.New()
	if err != nil {
		panic(err)
	}

	// デフォルトのwindowサイズとタイトルを変更する
	window.WindowOptions.Width = 1200
	window.WindowOptions.Height = 980
	window.WindowOptions.Title = "Gotron"

	// browser window を起動
	// window objectを使ってElectronBrowserWindow をコントロールするために
	// websocketを使って go と nodejsのブリッジを確立するよう
	done, err := window.Start()
	if err != nil {
		panic(err)
	}

	// dev toolsを使う場合は下記をコメントアウトする(せっかくなのでdevtools使ってみる)
	window.OpenDevTools()

	// アプリケーションがcloseするのを待つ
	<-done
}
