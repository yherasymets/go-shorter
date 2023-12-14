<p align="center" ><img src="https://media.tenor.com/L7YcQoDk9dsAAAAi/cut-it-out-scissors.gif" height="50"/></p> 
<h2 align="center">GO-SHORTER</h2>

&nbsp;&nbsp;&nbsp;&nbsp;This pet-project is a backend application written in the Go programming language, use gRPC for efficient communication between client and service. The project use PostgreSQL as the primary database for persistent data storage and Redis as in-memory cache for faster access to frequently used information. Also a web client was created with a simple UI, using the stdlib.

<h3>&nbsp;&nbsp;&nbsp;How to run application</h3>

For start web-client locally use Makefile and execute commands bellow:

- set environment variables: `make set-env`
- build client: `make build-client`
- start client: `make run-client`

For start server in Docker conteiner use docker-compose and Dockerfile:

- start go-shorter server, run command: `docker compose up` 