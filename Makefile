.PHONY: build start

build:
	GOOS=linux GOARCH=amd64 go build -o build/app/ad-service cmd/*.go

start: build
	go run cmd/*.go
