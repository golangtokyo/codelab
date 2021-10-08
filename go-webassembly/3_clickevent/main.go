package main

import "syscall/js"

func main() {
	// グローバルオブジェクト（Webブラウザはwindow）の取得
	window := js.Global()

	// window.document.getElementById("message")を実行
	message := window.Get("document").Call("getElementById", "message")

	// イベントリスナーとして登録するコールバックを作成
	cb := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		// HTMLを変更する
		message.Set("innerHTML", "Clicked!!")
		return nil
	})

	// イベントリスナーの登録
	// message.addEventListener("click", cb)
	message.Call("addEventListener", "click", cb)

	// Goのプログラムを終了させない
	select {}
}
