package main

import "syscall/js"

func main() {
	// グローバルオブジェクト（Webブラウザはwindow）の取得
	window := js.Global()

	// window.document.getElementById("message")を実行
	message := window.Get("document").Call("getElementById", "message")

	// HTMLを変更する
	message.Set("innerHTML", "Hello, WebAssembly")
}
