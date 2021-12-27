.PHONY: deps build up start debug-start down clean cleanall tidy test $(TOOLS) tools mocks lint gen-proto

SERVICE_NAME ?= ad-service
BINDIR ?= build/app

TOOLS += github.com/maxbrunsfeld/counterfeiter/v6

APP_CONTAINER_NAME ?= $(SERVICE_NAME)-app
DB_CONTAINER_NAME ?= $(SERVICE_NAME)-db-pg

INFO ?= [MAKE INFO]:
ERROR ?= [MAKE ERROR]:

deps:
	$(info $(INFO) download dependency packages)
	GO111MODULE=on go mod download

build: deps
	$(info $(INFO) build binary file and docker containers)
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o $(BINDIR)/$(SERVICE_NAME) cmd/ad/*.go
	docker-compose build

up:
	$(info $(INFO) starting application...)
	docker-compose up

start: build up

debug-start:
	LOG_LEVEL=debug make start

down:
	docker-compose down

clean:
	rm $(BINDIR)/$(SERVICE_NAME)
	docker stop $(APP_CONTAINER_NAME) $(DB_CONTAINER_NAME)
	docker rm $(APP_CONTAINER_NAME) $(DB_CONTAINER_NAME)

cleanall: clean
	docker rmi $(APP_CONTAINER_NAME) $(DB_CONTAINER_NAME)

tidy:
	go mod tidy

test:
	go test -race ./...

$(TOOLS): %:
	GOBIN=$(GOBIN) go install $*

tools: deps $(TOOLS)

mocks:
	@go generate ./...

lint:
	golangci-lint run ./...

gen-proto:
	protoc --proto_path=proto \
	--go_out=svc-ads \
	--go-grpc_out=. \
	svc_ads.proto
