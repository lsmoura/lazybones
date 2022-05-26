build-server:
	go build -o out/server cmd/server/*.go

build-web:
	GOARCH=wasm GOOS=js go build -o out/static/lib.wasm cmd/web/*.go
	cp "$(shell go env GOROOT)/misc/wasm/wasm_exec.js" out/static/
	cp cmd/web/index.html out/static/

build-all: build-web build-server

start: build-all
	cd out && ./server