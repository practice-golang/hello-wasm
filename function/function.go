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

func memberVar(this js.Value, d []js.Value) any {
	if len(d) < 2 {
		return "undefined"
	}

	objectName := d[0].String()
	memberName := d[1].String()

	f := g.Call("eval", objectName+"."+memberName)

	return f
}

func main() {
	c := make(chan struct{}, 0)
	g.Set("hello", js.FuncOf(Hello))
	g.Set("sum", js.FuncOf(Sum))
	g.Set("membervar", js.FuncOf(memberVar))
	<-c
}
