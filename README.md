<p align="center" ><img src="https://media.tenor.com/L7YcQoDk9dsAAAAi/cut-it-out-scissors.gif" height="50"/></p> 
<h2 align="center">GO-SHORTER</h2>

&nbsp;&nbsp;&nbsp;&nbsp;This pet-project is a backend application written in the Go programming language, use gRPC for efficient communication between client and service. The project use PostgreSQL as the primary database for persistent data storage and Redis as in-memory cache for faster access to frequently used information. Also a web client was created with a simple UI, using the stdlib.

-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------

<h4>How to run application</h4>

For start web-client locally in project root directory execute commands bellow:

- set environment variables: `source .env`
- build client: `make build-client`
- build server: `make build-server`
- start client: `make run-client`
- start server: `make run-server`

For start server in Docker conteiner use docker-compose and Dockerfile:

- start go-shorter server app, run command: `docker compose up`

Screenshots:

![screen-1](https://github.com/yherasymets/go-shorter/blob/main/screenshots/Screenshot1.png) 
![screen-2](https://github.com/yherasymets/go-shorter/blob/main/screenshots/Screenshot2.png) 
![screen-3](https://github.com/yherasymets/go-shorter/blob/main/screenshots/Screenshot3.png) 
![screen-4](https://github.com/yherasymets/go-shorter/blob/main/screenshots/Screenshot4.png) 
