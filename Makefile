dsn := $(shell cat .dsn)

.PHONY: db-down db-up run build clean up down bench

.ONESHELL:
SHELL := /bin/bash

help:
	@echo "Commands:"
	@echo " make run - Run the InOrder server"
	@echo " make build - Build the InOrder server into a binary"
	@echo " make db-down - Run database migrations down"
	@echo " make db-up - Run database migrations up"
	@echo " make test - Run unit tests"
	@echo " make clean - Clean up builds"
	@echo " make quickstart - Start the server with dummy data (0 Configuration)"
	@echo " make up - Start the server without dummy data"
	@echo " make down - Stop the server"
	@echo " make bench - Run benchmarks using apache workbench on GET: /api/items && POST: /api/orders"

clean:
	rm -f inorder

quickstart:
	touch config.yaml
	cat sample.config.yaml > config.yaml
	DB_VOLUME=./database/mysql-init/dump.sql:/docker-entrypoint-initdb.d/dump.sql docker compose up -d --build
up:
	DB_VOLUME=./database/migrations:/docker-entrypoint-initdb.d docker compose up -d --build
down:
	docker compose down

db-down:
	migrate -path database/migrations/ -database "mysql://${dsn}" -verbose down

db-up:
	migrate -path database/migrations/ -database "mysql://${dsn}" -verbose up

run:
	@echo ""
	@cat logo.txt
	@echo ""

	go run -ldflags="-s -w" cmd/main.go

build:
	go build -o inorder -ldflags="-s -w" cmd/main.go

test:
	@env INORDER_CONFIG=../../config.yaml go test -v ./...

bench:
	$(MAKE) quickstart
	@echo "Waiting for 10 seconds before starting benchmark"
	@sleep 10
	@echo "Logging in to get AuthToken"
	@AUTH=$$(curl -s -X POST http://localhost:4000/api/auth/login -H "Content-Type: application/json" -d '{"username":"admin", "password":"admin"}' | jq -r '.authToken');
	@echo "Extracted AuthToken: $${AUTH}";
	@echo ""
	@echo "Running Apache Benchmark on GET /api/items"
	ab -n 100000 -c 1000 -H "Authorization: Bearer $$AUTH" http://localhost:4000/api/items;
	@echo ""
	@echo "Waiting for 10 seconds before starting next benchmark"
	@sleep 10;
	@echo ""
	@echo "Running Apache Benchmark on POST /api/orders";
	@echo ""
	ab -l -n 100000 -c 1000 -p ./benchmark/create_order_body.json -T application/json  -H "Authorization: Bearer $$AUTH" http://localhost:4000/api/orders
	$(MAKE) down