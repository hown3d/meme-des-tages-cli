.PHONY: wasm
wasm:
	GOOS=js GOARCH=wasm go build -o  assets/main.wasm ./cmd/wasm/main.go
	cp "$(shell go env GOROOT)/misc/wasm/wasm_exec.js" $(shell pwd)/assets/

serve: wasm
	go run ./cmd/fileserver/main.go