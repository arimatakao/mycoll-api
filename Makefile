all:
	go run ./cmd/apiserver/main.go
	
build:
	go build ./cmd/apiserver/main.go

help:
	go run ./cmd/apiserver/main.go -help