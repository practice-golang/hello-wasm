package main // import "function.wasm"

import (
	"regexp"
	"syscall/js"
)

var g = js.Global()
var reFuncExist = regexp.MustCompile(`[(|)]+`)

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
		return js.Undefined()
	}

	varName := d[0].String()
	matches := reFuncExist.FindStringSubmatch(varName)

	if len(matches) > 0 {
		return js.Undefined()
	}

	// f := g.Call("eval", varName+"")
	f1 := g.Call("Function", "'use strict'; return "+varName)
	f2 := f1.Invoke()

	return f2
}

func callFunction(this js.Value, d []js.Value) any {
	if len(d) < 1 {
		return "undefined"
	}

	cmdString := d[0].String()
	matches := reFuncExist.FindStringSubmatch(cmdString)

	if len(matches) == 0 {
		return js.Undefined()
	}

	f1 := g.Call("Function", "'use strict'; return "+cmdString)
	f2 := f1.Invoke()

	return f2
}

func checkHost(this js.Value, d []js.Value) any {
	host := g.Get("window").Get("location").Get("host")

	result := js.Undefined()
	if host.String() == "localhost" {
		result = host
	}

	return result
}

func main() {
	c := make(chan struct{}, 0)
	g.Set("hello", js.FuncOf(Hello))
	g.Set("sum", js.FuncOf(Sum))
	g.Set("getvalue", js.FuncOf(getValue))
	g.Set("callfunction", js.FuncOf(callFunction))
	g.Set("checkHost", js.FuncOf(checkHost))
	<-c
}
