package main // import "hello.wasm"

import (
	"fmt"
	"syscall/js"
)

var g = js.Global()

func Hi(this js.Value, d []js.Value) any {
	name := "anonymous"
	if len(d) > 0 {
		name = d[0].String()
	}
	result := "Hi " + name + "!!"

	return result
}

func main() {
	fmt.Println("Hello, WebAssembly!")
	fmt.Println("This is fmt.Println from hello.wasm")

	c := make(chan struct{}, 0)
	g.Set("hi", js.FuncOf(Hi))
	<-c
}
