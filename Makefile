# set default shell
SHELL = bash -e -o pipefail

# variables
VERSION                  ?= $(shell cat ./VERSION)

# docker related
HOST_PORT				 ?= 8000
DOCKER_PORT				 ?= 8000
MYSQL_HOST_PORT		     ?= 3306
MYSQL_DOCKER_PORT		 ?= 3306
DOCKER_EXTRA_ARGS        ?=
DOCKER_REGISTRY          ?=
DOCKER_REPOSITORY        ?=
DOCKER_TAG               ?= ${VERSION}
DOCKER_BUILD_ARGS        ?=${DOCKER_EXTRA_ARGS} --build-arg version="${VERSION}"
IMAGE_NAME				 ?=kbaluyot/gomora

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

.PHONY:	run-windows
run-windows:	build
	./bin/gomora.exe

.PHONY:	build-docker
build-docker:
	docker build ${DOCKER_BUILD_ARGS} -t ${IMAGE_NAME}:${VERSION} -t ${IMAGE_NAME}:latest .

.PHONY: up
up:
	docker run -p ${HOST_PORT}:${DOCKER_PORT} -p ${MYSQL_HOST_PORT}:${MYSQL_DOCKER_PORT} ${IMAGE_NAME}:${VERSION}

proto-record:
	protoc --go_out=plugins=grpc:. --go_opt=paths=source_relative module/record/interfaces/http/grpc/pb/record.proto