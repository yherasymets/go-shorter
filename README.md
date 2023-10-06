`export PATH=$PATH:$(go env GOPATH)/bin`

## How to generate proto file:

inside /proto folder run command

`$ protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative *.proto`

## Run PostgreSQL in docker container:

`docker run --name postgres_db -e POSTGRES_PASSWORD=postgres -e POSTGRES_USER=postgres -d -p 8888:5432 postgres`

## Run Redis Server + Redis Insight in docker container:

`docker run -d --name redis-stack -p 6379:6379 -p 8001:8001 redis/redis-stack:latest`

## Migrations:

1. `$go install github.com/gobuffalo/pop/v6/soda@latest"`

2. `create database.yml config file`

## generate migration

3. `$(go env GOPATH)/bin/soda generate sql ./migrations/name_of_migration`

## run migration

4. `$(go env GOPATH)/bin/soda migrate up`