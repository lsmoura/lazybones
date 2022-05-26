FROM go:1.18-alpine3.16 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# build web
RUN GOARCH=wasm GOOS=js go build -o out/static/lib.wasm cmd/web/*.go && \
	cp "$(shell go env GOROOT)/misc/wasm/wasm_exec.js" out/static/ && \
	cp cmd/web/index.html out/static/

# build server
RUN go build -o out/server cmd/server/*.go

FROM alpine:3.16

WORKDIR /app
COPY --from=builder /app/out/* ./

EXPOSE 8080

CMD ["/app/server"]