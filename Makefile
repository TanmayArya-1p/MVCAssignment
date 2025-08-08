dsn := $(shell cat .dsn)


help:
	@echo "Commands:"
	@echo " make run - Run the InOrder server"
	@echo " make build - Build the InOrder server into a binary"
	@echo " make db-down - Run database migrations down"
	@echo " make db-up - Run database migrations up"
	@echo " make test - Run unit tests"

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

.PHONY: db-down db-up run build
