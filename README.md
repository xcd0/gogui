# gogui golangでのGUIアプリケーション作成テスト

GUIとして、Webベースで作るか、pureGoで作るか、検討中。
Golangのパッケージがいくつかあるのでとりあえず試す。

## 所感

クロスコンパイルを考えるとwebベースになる。
Cgoを使っているとgccgoを使う感じになる。
gccgoはガベージコレクションがよくないらしく長時間起動するアプリにはよくないらしい。
よってwebベースになる感じ。

Electronはメモリを200MBくらい使うのでちょっと微妙。
単純な実装がよさそう。

webviewがお手軽な感じ。

## Webベース

* [webview](https://github.com/zserge/webview)
	* とりあえず簡単そう
	* スクラッチから書き直されている最中らしい...

* [lorca](github.com/zserge/lorca)
	* PCにインストールされているChromeを使って動くらしい。
	* 外観はHTML5



### wasmベース

* [app](https://github.com/maxence-charriere/app)
	* 最近もちゃんと更新している。
	* サンプルが動かない。

* [vugu](https://github.com/vugu/vugu)
	* [vugu](https://www.vugu.org/)
	* まだexperimental
	* もりもり？な空気


### Electronベース

* [gotron](https://github.com/Equanox/gotron)
	* シンプル？な感じ
	* ちょっとメモリ消費低い感じがする。60MBとか。

* [go-astilectron](https://github.com/asticode/go-astilectron)
	* 名前が読みずらい、長い。その弊害がやたら多い気がする...

## CGo 使っててクロスコンパイルしずらいやつ

見た目、機能はよい。だが...\_(┐「ε:)\_

* [fyne](https://github.com/fyne-io/fyne)
	* 公式ビルドのGolang(gc)だとビルドできない（windows版）
	* gccgoでビルドしないといけない。ひどい\_(:3 」∠ )\_。
		* gccgoはガベージコレクションが微妙なので選択肢としてない。
	* クロスコンパイルはちょっとめんどそう。
		* cgo使ってるせいっぽい。
	* 見た目はよい。使い勝手も悪くなさそう。
	* OpenGL使えそう。

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

