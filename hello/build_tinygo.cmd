go env -w GOOS=js GOARCH=wasm
tinygo build -o ../hello_tinygo.wasm -target wasm
