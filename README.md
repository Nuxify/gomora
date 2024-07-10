# Gomora
A progressive framework-agnostic API template following CLEAN architecture and SOLID principles. DDD inspired :)

## Introduction

Gomora provides the example for a module-based gRPC and REST server suitable for building progressive APIs (from monolith to distributed microservices).

![image](https://github.com/kabaluyot/gomora/assets/38805756/b6d30884-15e0-4b2f-99bf-fad48bb6da87)

<img width="1419" alt="Screen Shot 2023-11-08 at 9 23 38 AM" src="https://github.com/kabaluyot/gomora/assets/38805756/4f7f77f5-b0a5-4911-b3fd-a1abcaca40bb">

<img width="1419" alt="Screen Shot 2023-11-08 at 9 23 52 AM" src="https://github.com/kabaluyot/gomora/assets/38805756/698b605b-63dd-4f83-ae52-06cd5eb3a157">

## Database Migration

Gomora uses go-migrate (https://github.com/golang-migrate/migrate) to handle migration. Download and change your migrate database command accordingly.

To create a schema, run:
- make schema NAME=<init_schema>

To migrate up, run:
- make migrate-up STEPS=<remove STEPS to apply all or specify step number>

To migrate down, run:
- make migrate-down STEPS=<remove STEPS to apply all or specify step number>

## Local Development

Setup the .env file first
- cp .env.example .env

To bootstrap everything, run:
- make

The command above will install, build, and run the binary

For manual install:
- make install

For lint:
- make lint

Just ensure you installed golangci-lint.

To test:
- make test

For manual build:
- make build
- NOTE: the output for this is in bin/

## Docker Build

To build, run:
- make run

To run the container:
- make up

## License

[MIT](https://choosealicense.com/licenses/mit/)

Made with ❤️ at [Nuxify](https://nuxify.tech)
