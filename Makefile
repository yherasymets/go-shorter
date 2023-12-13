.PHONY:
.SILENT:

clear:
	rm -rf ./bin/*
	go clean

build-client:
	go build -o ./bin/client ./cmd/client

build-server:
	go build -o ./bin/server ./cmd/server

# set environment variables
set-env:
	source .env

run-client:
	./bin/client

run-server:
	./bin/server

# build docker image 
build-image:
	docker build -t go-shorter:v0.0.1 .

# start container with app
start-container:
	docker run --name url-shorter -p 8081:8081 go-shorter:v0.0.1