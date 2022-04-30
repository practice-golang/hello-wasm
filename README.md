# Taste WASM

## First
* Compile and Run `server/server.go`
* Open localhost:80 in web browser
* Open `Developer console` in web browser

## Dir

### wasm
* hello: Just Println
* function: Use func

### Simple server
* server: server for wasm

## My preference

### vscode

```json
"go.toolsEnvVars": {
    "GOARCH": "wasm",
    "GOOS": "js",
    "GOBIN": "${workspaceFolder}/bin"
}
```

## Requirement to use tinygo
* `Go` >= 1.18 which support latest tinygo - https://go.dev/dl
* Tinygo - https://jawabsoal.live/baca-https-github.com/tinygo-org/tinygo/releases
* LLVM, MinGW - https://github.com/brechtsanders/winlibs_mingw/releases
* Binaryen (Need only wasm-opt.exe) - https://github.com/WebAssembly/binaryen/releases
* wasm_exec_tinygo.js
    * Not compatible with `{go_root}/misc/wasm_exec.js` of Go so, have to use `{tinygo_root}/targets/wasm_exec.js`
    * `wasm_exec_tinygo.js:313` - Because of [`syscall/js.finalizeRef`](https://github.com/tinygo-org/tinygo/issues/1140), commented it
