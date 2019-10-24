# fyne

## intro
go get -u -v fyne.io/fyne

* 公式ビルドのGolang(gc)だとビルドできない（windows版）
* gccgoでビルドしないといけない
	* MSYSではなくMinGW64を起動
	* `pacman -S mingw-w64-x86_64-gcc`
	* `go get -u -v fyne.io/fyne/cmd/fyne_demo`
	* `/c/Users/username/go/bin/fine_demo.exe`

	で動く。
* クロスコンパイルはちょっとめんどそう。
	* cgo使ってるせいっぽい。
* 見た目はよい。使い勝手も悪くなさそう。
* OpenGL使えそう。
