package main

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/yherasymets/go-shorter/internal/shorter"
	"github.com/yherasymets/go-shorter/proto"
	"github.com/yherasymets/go-shorter/repo"
	"google.golang.org/grpc"
)

var gRPCport = os.Getenv("GRPC_PORT")

func main() {
	listener, err := net.Listen("tcp", gRPCport)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	db := repo.Connection()
	cache := repo.RedisCache()
	defer cache.Close()
	log.Printf("starting redis client: %v", cache.Ping(context.Background()).Val())

	service := grpc.NewServer()
	server := &shorter.GRPCServer{
		DB:    db,
		Cache: cache,
	}
	proto.RegisterShorterServer(service, server)

	log.Printf("starting gRPC listener on port %v", gRPCport)
	if err := service.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
