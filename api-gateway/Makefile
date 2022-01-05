.PHONY: deps build clean tidy test

SERVICE_NAME ?= api-gateway
BIN_DIR ?= build

deps:
	GO111MODULE=on go mod download

build: deps
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o $(BIN_DIR)/$(SERVICE_NAME) *.go

clean:
	rm $(BIN_DIR)/$(SERVICE_NAME)

tidy:
	go mod tidy

test:
	go test -race ./...
