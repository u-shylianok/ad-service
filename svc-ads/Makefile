.PHONY: deps build clean tidy test gen_mocks gen_proto

SERVICE_NAME ?= svc-ads
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

test:
	go test -race ./...

gen_mocks:
	@go generate ./...

gen_proto:
	protoc -I=. -I=../ \
	--go_out=. --go_opt=module=$(MODULE) \
	--go-grpc_out=. --go-grpc_opt=module=$(MODULE) \
	./proto/*.proto

lint:
	golangci-lint run ./...
