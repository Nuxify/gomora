include .env

# set default shell
SHELL = bash -e -o pipefail

default: run

.PHONY:	install
install:
	go mod tidy
	go mod vendor

.PHONY:	lint
lint:
	golangci-lint run 

.PHONY:	build
build:
	mkdir -p bin
	go build -race -o bin/gomora \
	    cmd/main.go

.PHONY:	test
test:
	go test -race -v -p 1 ./...
	
.PHONY:	run
run:	build
	./bin/gomora

.PHONY: up
up:
	docker compose down
	docker compose up -d --build

.PHONY:	schema
schema:
	mkdir -p db/migration
	migrate create -ext sql -dir db/migration -seq ${NAME}

.PHONY: migrate-up
migrate-up: 
	migrate -path db/migration/ -database "mysql://${DB_USERNAME}:${DB_PASSWORD}@tcp(${DB_HOST}:${DB_PORT})/${DB_DATABASE}" -verbose up ${STEPS}

.PHONY: migrate-down
migrate-down:
	migrate -path db/migration/ -database "mysql://${DB_USERNAME}:${DB_PASSWORD}@tcp(${DB_HOST}:${DB_PORT})/${DB_DATABASE}" -verbose down ${STEPS}

proto-record:
	protoc --go_out=plugins=grpc:. --go_opt=paths=source_relative module/record/interfaces/http/grpc/pb/record.proto