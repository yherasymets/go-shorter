## How to generate proto file:

Inside `./proto` folder run command:

`$ protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative *.proto`

## Run PostgreSQL in docker container:

`$ docker run --name postgres_db -e POSTGRES_PASSWORD=postgres -e POSTGRES_USER=postgres -d -p 5433:5432 postgres`

## Run Redis Server + Redis Insight in docker container:

`$ docker run -d --name redis-stack -p 6379:6379 -p 8001:8001 redis/redis-stack:latest`
