# dealls-recruitment-dating-app-api

Dealls dev recruitment test 2024: Dating App API.

## Development Requirements

- [Go](https://go.dev/dl/) v1.23.2+
- [Docker Desktop](https://docs.docker.com/desktop/release-notes/) v4.34.2+

## Setup

Create configuration files:

- Copy `.env.example` to `.env` and `.env.testing`. Set DB user and password. You might also want to use different DB for testing.
- Copy `docker/postgres.cfg.example` to `docker/postgres.cfg`. Set DB user and password.

Download/install dependencies:

```sh
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

```sh
rel migrate -dir=rel/migrations
rel migrate -dir=rel/migrations -dsn=<your testing DATABASE_URL>
```

## Run app with live reload

```sh
air
```

## Run tests on API handlers

**Note:** You don't need to run app before running the tests.

```sh
go test -v -p=1 ./internal/handlers/... -args -envfile=$(pwd)/.env.testing
```

## Project directory structure

```
.
├── cmd                                           # Entry points (main function), to execute any initial setups and the main loop.
│   └── dating-app-api                            # We can have more than one entry points for different use cases.
│       └── main.go
├── configs                                       # Single point of truth (SPOT) for app configurations, such as environment variables.
│   └── env
│       └── env.go
├── docker                                        # Configurations for Docker containers.
│   ├── postgres.cfg
│   └── postgres.cfg.example
├── internal                                      # Internal packages, not for import by other projects.
│   ├── crypto                                    # All about cryptography.
│   │   ├── jwtutil                               # JSON Web Token (JWT).
│   │   │   └── jwtutil.go
│   │   ├── pwdutil                               # Password hashing.
│   │   │   └── pwdutil.go
│   │   └── signingkey                            # Key management, for signing and verifying JWTs.
│   │       └── signingkey.go
│   ├── db                                        # Database connection.
│   │   ├── models                                # REL models, represent DB tables.
│   │   │   ├── User.go
│   │   │   └── UserProfile.go
│   │   └── client.go
│   ├── handlers                                  # API handlers.
│   │   ├── access                                # Handlers for user login.
│   │   │   ├── accessform
│   │   │   │   └── form.go
│   │   │   ├── access.go
│   │   │   └── access_test.go
│   │   ├── authorization                         # Authorization middleware. It decodes and verifies JWTs.
│   │   │   └── authorization.go
│   │   ├── myprofile                             # Handlers for showing and updating current user profile.
│   │   │   ├── myprofileform
│   │   │   │   └── form.go
│   │   │   ├── myprofile.go
│   │   │   └── myprofile_test.go
│   │   └── registration                          # Handlers for user registration.
│   │       ├── registrationform
│   │       │   └── form.go
│   │       ├── registration.go
│   │       └── registration_test.go
│   ├── helpers                                   # For some shared functions. However, it's better to have more proper grouping, like crypto.
│   │   ├── jsonresponse                          # JSON response formatting.
│   │   │   └── jsonresponse.go
│   │   └── testinghelper                         # Testing helper.
│   │       ├── setup.go                          # Test setup.
│   │       ├── seeds.go                          # Data seeds. We can put it somewhere else if needed.
│   │       └── response.go                       # API response checkers.
│   └── routers                                   # Manage API routes.
│       └── v1                                    # API routes version 1.
│           └── v1-router.go
├── rel                                           # Since we use REL CLI to manage DB migrations,
│   └── migrations                                # we need to put the migration codes outside of internal packages.
│       ├── 20241102141500_create_users.go
│       └── ...
├── .air.toml
├── docker-compose.yaml
├── .env
├── .env.example
├── .env.testing
├── .gitignore
├── go.mod
├── go.sum
├── LICENSE
└── README.md
```
