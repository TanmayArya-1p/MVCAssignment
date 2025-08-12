dsn := $(shell cat .dsn)

.PHONY: db-down db-up run build clean up

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
	@echo "\n"
	@cat logo.txt
	@echo "\n"

	go run cmd/main.go

build:
	go build -o inorder cmd/main.go

test:
	@env INORDER_CONFIG=../../config.yaml go test -v ./...
