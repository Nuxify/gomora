# Gomora

A progressive framework-agnostic API template following CLEAN architecture and SOLID principles. DDD inspired :)

## Introduction

Gomora provides the example for a module-based gRPC and REST server suitable for building progressive APIs (from monolith to distributed microservices).

![image](https://github.com/kabaluyot/gomora/assets/38805756/b6d30884-15e0-4b2f-99bf-fad48bb6da87)

<img width="1419" alt="Screen Shot 2023-11-08 at 9 23 38 AM" src="https://github.com/kabaluyot/gomora/assets/38805756/4f7f77f5-b0a5-4911-b3fd-a1abcaca40bb">

<img width="1419" alt="Screen Shot 2023-11-08 at 9 23 52 AM" src="https://github.com/kabaluyot/gomora/assets/38805756/698b605b-63dd-4f83-ae52-06cd5eb3a157">

## Local Development

Setup the .env file first

```bash
cp .env.example .env
```

To bootstrap everything, run:

```bash
make
```

The command above will install, build, and run the binary

For manual install:

```bash
make install
```

For lint:

```bash
make lint
```

Just ensure you installed golangci-lint.

To test:

```bash
make test
```

For manual build:

```bash
make build

# The output for this is in bin/
```

## Docker Build

To build, run:

```bash
make run
```

To run the container:

```bash
make up
```

## Database Migration

Gomora uses go-migrate (https://github.com/golang-migrate/migrate) to handle migration. Download and change your migrate database command accordingly.

To create a schema, run:

```bash
make schema NAME=<init_schema>
```

To migrate up, run:

```bash
STEPS=<remove STEPS to apply all or specify step number> make migrate-up
```

To migrate down, run:

```bash
STEPS=<remove STEPS to apply all or specify step number> make migrate-down
```

To check migrate version, run:

```bash
make migrate-version
```

To force migrate, run:

```bash
STEPS=<specify step number> make migrate-force
```

## License

[MIT](https://choosealicense.com/licenses/mit/)

Made with ❤️ at [Nuxify](https://nuxify.tech)
