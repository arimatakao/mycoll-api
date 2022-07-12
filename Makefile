all:
	go run ./cmd/apiserver/main.go -config-path configs/config.toml
	
build:
	go build ./cmd/apiserver/main.go

help:
	go run ./cmd/apiserver/main.go -help