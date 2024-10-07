# Gomora

A progressive framework-agnostic API template following CLEAN architecture and SOLID principles. DDD inspired :)

## Introduction

Gomora provides the example for a module-based gRPC and REST server suitable for building progressive APIs (from monolith to distributed microservices).

<img width="1416" alt="Screen Shot 2024-10-07 at 9 43 31 AM" src="https://github.com/user-attachments/assets/736b9813-f086-4fd4-aeb3-98e08df51e7e">

<img width="1416" alt="Screen Shot 2024-10-07 at 9 41 10 AM" src="https://github.com/user-attachments/assets/0225919b-53d1-4900-9b9e-8d5d5f616d73">

<img width="1416" alt="Screen Shot 2024-10-07 at 9 44 17 AM" src="https://github.com/user-attachments/assets/a5871a22-be66-4236-8635-6166624249c1">

<img width="1416" alt="Screen Shot 2024-10-07 at 9 46 05 AM" src="https://github.com/user-attachments/assets/578de61d-73cb-47b9-8806-ed58a7b281a1">

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
