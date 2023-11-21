package main

import (
	"context"
	"net"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/yherasymets/go-shorter/internal/shorter"
	"github.com/yherasymets/go-shorter/proto"
	"github.com/yherasymets/go-shorter/repo"
	"google.golang.org/grpc"
)

var gRPCport = os.Getenv("GRPC_PORT")

func main() {
	listener, err := net.Listen("tcp", gRPCport)
	if err != nil {
		logrus.Errorf("failed to listen: %v", err)
		return
	}

	db := repo.Connection()
	cache := repo.RedisCache()
	defer cache.Close()
	logrus.Infof("starting redis client: %v", cache.Ping(context.Background()).Val())

	service := grpc.NewServer()
	server := &shorter.GRPCServer{
		DB:    db,
		Cache: cache,
	}
	proto.RegisterShorterServer(service, server)

	logrus.Infof("starting gRPC listener on port " + gRPCport)
	if err := service.Serve(listener); err != nil {
		logrus.Fatal(err)
	}
}
