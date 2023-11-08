package main

import (
	"context"
	"net"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/yherasymets/go-shorter/pkg/shorter"
	"github.com/yherasymets/go-shorter/proto"
	"github.com/yherasymets/go-shorter/repo"
	"google.golang.org/grpc"
)

var gRPCport = os.Getenv("GRPC_PORT")

func main() {
	l, err := net.Listen("tcp", gRPCport)
	if err != nil {
		logrus.Errorf("failed to listen: %v", err)
		return
	}
	db := repo.Conn()
	cache := repo.RedisClient()
	defer cache.Close()
	logrus.Infof("starting redis client: %v", cache.Ping(context.Background()).Val())

	service := grpc.NewServer()
	server := &shorter.GRPCServer{
		DB:    db,
		Cache: cache,
	}
	proto.RegisterShorterServer(service, server)

	logrus.Infof("starting gRPC listener on port " + gRPCport)
	if err := service.Serve(l); err != nil {
		logrus.Fatal(err)
	}
}
