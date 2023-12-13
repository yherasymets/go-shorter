FROM golang:1.21.4-alpine3.18 AS builder

COPY . /github.com/yherasymets/go-shorter/
WORKDIR /github.com/yherasymets/go-shorter/

RUN go mod download
RUN go build -o ./bin/go-shorter-server cmd/server/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=0 /github.com/yherasymets/go-shorter/bin/go-shorter-server .
COPY --from=0 /github.com/yherasymets/go-shorter/.env .

EXPOSE 8081

RUN source .env
CMD ["./go-shorter-server"]