# gogui golangでのGUIアプリケーション作成テスト

GUIとして、Webベースで作るか、pureGoで作るか、検討中。
Golangのパッケージがいくつかあるのでとりあえず試す。

方向性として
* cgoのいるものはアウト。
* windows,mac,linuxで動くこと。


## 所感

クロスコンパイルを考えるとwebベースになる。
デスクトップベースではGCCに依存する感じ、よくない。
よってwebベースになる感じ。

Electronはメモリを200MBくらい使うのでちょっと微妙。
単純な実装がよさそう。

webviewがお手軽な感じ。ただいまは作り直されているタイミングなので時期尚早。

pureGoのいい感じなGUIライブラリはよ...（自分でつくれ

## Webベース

* [webview](https://github.com/zserge/webview) star 5.3k (2020/01/27)
	* とりあえず簡単そう
	* 一番楽かもしれない
	* サンプル動かしてwindowsで15MBくらいでメモリが20MB弱
	* スクラッチから書き直されている最中らしい...
		* 「I expect to finish the new branch before March 2020, but no hard deadlines.」
	* 作り直しが終わってからかなあ...
	
	* 中国語の情報でそれをgoogle翻訳にかけたものですが、
	[「以前はデスクトップソフトウェアとしてSciterを学びたかったのですが、  
	サポートされているcssレベルが低すぎて、ボタンの背景色を変更できず、悲しかったです。  
	後でネットワークを調べて、Lorca webviewの 2つの興味深いことがわかりました。  
	SciterよりもHTMLおよびCSSのサポートがよいです。」](https://memory6181.top/post/2019-05-22/)  
	とのことで、go-sciterよりはよさそう。
		* まあ更新されているかもだけれども

* [lorca](github.com/zserge/lorca) star 5k (2020/01/27)
	* PCにインストールされているChromeを使って動くらしい。
	* 外観はHTML5

## OpenGL ESベース

* [Gio](https://git.sr.ht/~eliasnaur/gio) githubではない
	* 良さげな感じはする。pureGoっぽい。
	* github使ってないので見方がわかりずらい。
	* Immediate mode GUI programs in Go for Android, iOS, macOS, Linux, FreeBSD, Windows, and WebAssembly
	* experimental
	* 作者はフルタイムで作っているらしい。

### wasmベース

* [app](https://github.com/maxence-charriere/app) star 3.3k (2020/01/27)
	* サンプルが動かない。
	* なんか微妙そう

* [vugu](https://github.com/vugu/vugu)
	* [vugu](https://www.vugu.org/)
	* まだexperimental
	* もりもり？な空気


### Electronベース

* [gotron](https://github.com/Equanox/gotron)
	* シンプル？な感じ
	* ちょっとメモリ消費低い感じがする。60MBとか。

* [go-astilectron](https://github.com/asticode/go-astilectron) star 3k (2020/01/27)
	* 名前が読みずらい、長い。その弊害がやたら多い気がする...

## CGo 使っててクロスコンパイルしずらいやつ

見た目、機能はよい。だが...\_(┐「ε:)\_

* [fyne](https://github.com/fyne-io/fyne) star 7.6k (2020/01/27)
	* 公式ビルドのGolang(gc)だとビルドできない（windows版）
	* gccgoでビルドしないといけない。ひどい\_(:3 」∠ )\_。
		* gccgoはガベージコレクションが微妙なので選択肢としてない。
	* クロスコンパイルはちょっとめんどそう。
		* cgo使ってるせいっぽい。
	* 見た目はよい。使い勝手も悪くなさそう。
	* OpenGL使えそう。

* [g3n](https://github.com/g3n/engine) star 928 (2020/01/27)
	* 公式ビルドのGolang(gc)だとビルドできない（windows版）
	* gccgoでビルドしないといけない。ひどい\_(:3 」∠ )\_。
	* OpenGLがちゃんと動いている。
	* 重いか軽いかは判断がつかない。
		* macで25fpsくらい

## 未調査

* [ui](https://github.com/andlabs/ui) star 7.3k (2020/01/27)
	* OS依存はlibuiで吸収。現実的ではあると思う。つまりcgo。
* [walk](https://github.com/lxn/walk) star 4.2k (2020/01/27)
	* windowsのみ\_(:3 」∠ )\_
* [gowd](https://github.com/dtylman/gowd)  star 236 (2020/01/27)
	* nwjs
* [go-sciter](https://github.com/sciter-sdk/go-sciter) star 1.7k (2020/01/27)
	* 一番サンプル画像がいっぱい張ってある。
	* でもこのサンプルはsciterのサンプルでgo-sciterではない。
	* ラッパーなのでシングルバイナリというわけにはいかなそう。

