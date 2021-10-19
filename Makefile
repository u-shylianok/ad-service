.PHONY: build start up down clean cleanall

build:
	GOOS=linux GOARCH=amd64 go build -o build/app/ad-service cmd/*.go
	docker-compose build

start: build
	docker-compose up

up:
	docker-compose up

down:
	docker-compose down

clean:
	rm build/app/ad-service
	docker stop ad-service-app ad-service-db-pg
	docker rm ad-service-app ad-service-db-pg

cleanall: clean
	docker rmi ad-service-app ad-service-db-pg