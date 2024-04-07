package main

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/yherasymets/go-shorter/internal/service"
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

	db := repo.NewRepository()
	cache := repo.NewCache()
	defer cache.Close()
	log.Printf("starting redis client: %v", cache.Ping(context.Background()).Val())

	gRPCserver := grpc.NewServer()
	shorterService := service.NewService(db, cache)
	proto.RegisterShorterServer(gRPCserver, shorterService)

	log.Printf("starting gRPC listener on port %v", gRPCport)
	if err := gRPCserver.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
