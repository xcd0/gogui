# gogui golangでのGUIアプリケーション作成テスト

GUIとして、Webベースで作るか、pureGoで作るか、検討中。
Golangのパッケージがいくつかあるのでとりあえず試す。

## Webベース

### wasmベース

* [webview](https://github.com/zserge/webview)

* [app](https://github.com/maxence-charriere/app)

* [vugu](https://github.com/vugu/vugu)
	* [vugu](https://www.vugu.org/)
	* まだexperimental


### Electronベース

* [gotron](https://github.com/Equanox/gotron)
	* シンプル？な感じ

* [go-astilectron](https://github.com/asticode/go-astilectron)
	* 名前が読みずらい、長い。その弊害がやたら多い気がする...

## cGo

* [fyne](https://github.com/fyne-io/fyne)
	* 公式ビルドのGolang(gc)だとビルドできない（windows版）
	* gccgoでビルドしないといけない。ひどい\_(:3 」∠ )\_。
		* gccgoはガベージコレクションが微妙なので選択肢としてない。
	* クロスコンパイルはちょっとめんどそう。
		* cgo使ってるせいっぽい。
	* 見た目はよい。使い勝手も悪くなさそう。
	* OpenGL使えそう。



## OpenGLベース？
* [g3n](https://github.com/g3n/engine)
	* 公式ビルドのGolang(gc)だとビルドできない（windows版）
	* gccgoでビルドしないといけない。ひどい\_(:3 」∠ )\_。
	* OpenGLがちゃんと動いている。
	* 重いか軽いかは判断がつかない。
		* macで25fpsくらい

## 未調査

* [ui](https://github.com/andlabs/ui)
* [walk](https://github.com/lxn/walk)
* [gowd](https://github.com/dtylman/gowd)
	* nwjs

