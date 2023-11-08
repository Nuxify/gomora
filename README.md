# Gomora
A progressive framework-agnostic API template following CLEAN architecture and SOLID principles. DDD inspired :)

![image](https://github.com/kabaluyot/gomora/assets/38805756/b6d30884-15e0-4b2f-99bf-fad48bb6da87)

## Introduction

Gomora provides the example for a module-based gRPC and REST server suitable for building progressive APIs (from monolith to distributed microservices).

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

## Contributors ✨
[![](https://avatars0.githubusercontent.com/u/38805756?s=90&u=96545a7174420f0ae00a9511c74e6ed74a9e5319&v=4)](https://github.com/kabaluyot)

Made with ❤️ at [Nuxify](https://nuxify.tech)
