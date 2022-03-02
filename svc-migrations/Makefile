.PHONY: deps build clean tidy

SERVICE_NAME ?= svc-migrations
MODULE ?= github.com/u-shylianok/ad-service/$(SERVICE_NAME)
BIN_DIR ?= build

deps:
	GO111MODULE=on go mod download

build: deps
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o $(BIN_DIR)/$(SERVICE_NAME) *.go

clean:
	rm $(BIN_DIR)/$(SERVICE_NAME)

tidy:
	go mod tidy
