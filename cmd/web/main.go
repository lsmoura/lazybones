//go:build js && wasm

package main

import (
	"syscall/js"
)

const message = "Hello, WebAssembly!"

func main() {
	document := js.Global().Get("document")
	h2 := document.Call("createElement", "h2")
	h2.Set("innerHTML", message)
	document.Get("body").Call("appendChild", h2)
}
