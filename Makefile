.PHONY: deps build up start down clean cleanall

SERVICE_NAME ?= ad-service
APP_CONTAINER_NAME ?= $(SERVICE_NAME)-app
DB_CONTAINER_NAME ?= $(SERVICE_NAME)-db-pg
BINDIR ?= build/app

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

down:
	docker-compose down

clean:
	rm $(BINDIR)/$(SERVICE_NAME)
	docker stop $(APP_CONTAINER_NAME) $(DB_CONTAINER_NAME)
	docker rm $(APP_CONTAINER_NAME) $(DB_CONTAINER_NAME)

cleanall: clean
	docker rmi $(APP_CONTAINER_NAME) $(DB_CONTAINER_NAME)