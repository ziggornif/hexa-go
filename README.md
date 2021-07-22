# hexa-go

<img src="./logo.png" alt="logo" width="100">

Go hexagonal architecture implementation example.

## Requirements

Before running this project locally, you need these dependencies
- Go 1.15+
- docker

## Database

```
```yaml
docker run --ulimit memlock=-1:-1 -it --rm=true --memory-swappiness=0 \
    --name postgres-hexago -e POSTGRES_USER=hexago \
    -e POSTGRES_PASSWORD=hexago -e POSTGRES_DB=hexa-go \
    -p 5432:5432 postgres:13.1
```

## Building the server

```
make build
```

## Running the server locally

Run the following command :

```
make run
```

Open browser to http://localhost:8080 to start.

## Testing and linting

Run linting task :

```
make lint
```

Run tests suites :

```
make test
```
