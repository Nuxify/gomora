# set default shell
SHELL = bash -e -o pipefail

# variables
VERSION                  ?= $(shell cat ./VERSION)

now=$(shell date +"%Y%m%d%H%M%S")

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

proto-record:
	protoc --go_out=plugins=grpc:. --go_opt=paths=source_relative module/record/interfaces/http/grpc/pb/record.proto