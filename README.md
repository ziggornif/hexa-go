# hexa-go

<img src="./logo.png" alt="logo" width="100">

Go hexagonal architecture implementation example.

## Requirements

Before running this project locally, you need these dependencies
- Go 1.15+
- docker

## Project architecture

### Handler directory

Presentation / exposition layer is available on the handlers directory.

This directory contains all endpoints like : 
- rest
- graphql
- web
- console
- ...

Each Handler is stored in a dedicated directory.

Ex: rest endpoint are stored in rest dir

### Package directory

A package is a set of files corresponding to a business or technical domain.

The packages directory contains all the domains directories .

Each domain directory contains the associated model, service and repository layers

Example :

The `todo` domain contains todo model definition, service (use cases) functions and repository (db access) function.

Each domain directory can be easily moved out of the project and turned into a Go module.

### Infra directory

The infra directory contains project infrastructure stuffs like
- server management
- logger
- configuration
- db connections
- metrics (open-telemetry)
- ...

## Build the server

```
make build
```

## Run database

```yaml
docker run --ulimit memlock=-1:-1 -it --rm=true --memory-swappiness=0 \
    --name postgres-hexago -e POSTGRES_USER=hexago \
    -e POSTGRES_PASSWORD=hexago -e POSTGRES_DB=hexa-go \
    -p 5432:5432 postgres:13.1
```

## Run the server locally

Run the following command :

```
make run
```

Open browser to http://localhost:8080 to start.

## Running the tests

```
make test
```

## Running linter

```
make lint
```