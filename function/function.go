package main // import "function.wasm"

import (
	"syscall/js"
)

var g = js.Global()

func Hello(this js.Value, d []js.Value) any {
	name := "anonymous"
	if len(d) > 0 {
		name = d[0].String()
	}
	result := "Hello " + name + "!!"

	return result
}

func Sum(this js.Value, d []js.Value) any {
	result := 0

	for i := 0; i < len(d); i++ {
		result += d[i].Int()
	}

	return result
}

func getValue(this js.Value, d []js.Value) any {
	if len(d) < 1 {
		return "undefined"
	}

	varName := d[0].String()

	// f := g.Call("eval", varName+"")
	f1 := g.Call("Function", "'use strict'; return "+varName)
	f2 := f1.Invoke()

	return f2
}

func main() {
	c := make(chan struct{}, 0)
	g.Set("hello", js.FuncOf(Hello))
	g.Set("sum", js.FuncOf(Sum))
	g.Set("getvalue", js.FuncOf(getValue))
	<-c
}
