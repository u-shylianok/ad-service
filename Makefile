.PHONY: build docker_build docker_up docker_up_mig start start_mig debug_start down clean cleanall tidy test $(TOOLS) tools gen_mocks gen_proto lint

TOOLS += github.com/maxbrunsfeld/counterfeiter/v6@latest
TOOLS += github.com/pressly/goose/v3/cmd/goose@latest

API_GATEWAY ?= api-gateway
SVC_ADS ?= svc-ads
SVC_AUTH ?= svc-auth
SVC_MIGRATIONS ?= svc-migrations
DB_ADS ?= db-ads
DB_AUTH ?= db-auth

CONTAINERS ?= $(API_GATEWAY) $(SVC_ADS) $(SVC_AUTH) $(DB_ADS) $(DB_AUTH) $(SVC_MIGRATIONS)

INFO ?= [MAKE INFO]:
ERROR ?= [MAKE ERROR]:

build:
	cd $(API_GATEWAY) && $(MAKE) --silent build
	cd $(SVC_ADS) && $(MAKE) --silent build
	cd $(SVC_AUTH) && $(MAKE) --silent build
	cd $(SVC_MIGRATIONS) && $(MAKE) --silent build

docker_build:
	docker-compose build

docker_up:
	docker-compose up

docker_up_mig:
	docker-compose --profile migrations up

start: build docker_build docker_up

start_mig: build docker_build docker_up_mig

debug_start:
	LOG_LEVEL=debug $(MAKE) start

down:
	docker-compose down

clean:
	cd $(API_GATEWAY) && $(MAKE) --silent clean
	cd $(SVC_ADS) && $(MAKE) --silent clean
	cd $(SVC_AUTH) && $(MAKE) --silent clean
	docker stop $(CONTAINERS)
	docker rm $(CONTAINERS)

cleanall: clean
	docker rmi $(CONTAINERS)

tidy:
	cd $(API_GATEWAY) && $(MAKE) --silent tidy
	cd $(SVC_ADS) && $(MAKE) --silent tidy
	cd $(SVC_AUTH) && $(MAKE) --silent tidy

test:
	cd $(API_GATEWAY) && $(MAKE) --silent test
	cd $(SVC_ADS) && $(MAKE) --silent test
	cd $(SVC_AUTH) && $(MAKE) --silent test

$(TOOLS): %:
	GOBIN=$(GOBIN) go install $*

tools: $(TOOLS)

gen_mocks:
	cd $(SVC_ADS) && $(MAKE) --silent gen_mocks
	cd $(SVC_AUTH) && $(MAKE) --silent gen_mocks

gen_proto:
	cd $(SVC_ADS) && $(MAKE) --silent gen_proto
	cd $(SVC_AUTH) && $(MAKE) --silent gen_proto

lint:
	cd $(API_GATEWAY) && golangci-lint run
	cd $(SVC_ADS) && golangci-lint run
	cd $(SVC_AUTH) && golangci-lint run
