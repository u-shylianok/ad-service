.PHONY: build up start debug-start down clean cleanall tidy test $(TOOLS) tools gen_mocks gen_proto lint

TOOLS += github.com/maxbrunsfeld/counterfeiter/v6

API_GATEWAY ?= api-gateway
SVC_ADS ?= svc-ads
SVC_AUTH ?= svc-auth
DB_ADS ?= db-ads
DB_AUTH ?= db-auth

CONTAINERS ?= $(API_GATEWAY) $(SVC_ADS) $(SVC_AUTH) $(DB_ADS) $(DB_AUTH)

INFO ?= [MAKE INFO]:
ERROR ?= [MAKE ERROR]:

build:
	cd $(API_GATEWAY) && $(MAKE) --silent build
	cd $(SVC_ADS) && $(MAKE) --silent build
	cd $(SVC_AUTH) && $(MAKE) --silent build
	docker-compose build

up:
	docker-compose up

start: build up

debug-start:
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

tools: deps $(TOOLS)

gen_mocks:
	cd $(SVC_ADS) && $(MAKE) --silent gen_mocks
	cd $(SVC_AUTH) && $(MAKE) --silent gen_mocks

gen_proto:
	cd $(SVC_ADS) && $(MAKE) --silent gen_proto
	cd $(SVC_AUTH) && $(MAKE) --silent gen_proto

lint:
	golangci-lint run ./$(SVC_ADS)/...
