# dealls-recruitment-dating-app-api

Dealls dev recruitment test 2024: Dating App API.

## Development Requirements

- [Go](https://go.dev/dl/) v1.23.2+
- [Docker Desktop](https://docs.docker.com/desktop/release-notes/) v4.34.2+

## Setup

Create configuration files:

- Copy `.env.example` to `.env`. Set DB user and password.
- Copy `docker/postgres.cfg.example` to `docker/postgres.cfg`. Set DB user and password.

Download install dependencies:

```
# Start dependency containers, like PostgreSQL database
docker compose -p dating-app up -d --remove-orphans

# Download Go module dependencies
go mod download

# Install rel (database migrator)
go install github.com/go-rel/cmd/rel@latest

# Install air (live reloader)
go install github.com/air-verse/air@latest
```

## Migrate database

```
rel migrate -dir=rel/migrations
```

## Run with live reload

```
air
```
