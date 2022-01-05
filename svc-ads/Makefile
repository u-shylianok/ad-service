.PHONY: deps build clean tidy test gen_mocks gen_proto

SERVICE_NAME ?= svc-ads
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
	protoc --go_out=client --go_opt=paths=import \
	--go-grpc_out=client --go-grpc_opt=paths=import \
	./proto/*.proto

lint:
	golangci-lint run ./...
