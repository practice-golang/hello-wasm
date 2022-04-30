go env -w GOOS=js GOARCH=wasm
tinygo build -o ../function_tinygo.wasm -target wasm
