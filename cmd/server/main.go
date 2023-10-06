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

func main() {
	l, err := net.Listen("tcp", os.Getenv("GRPC_PORT"))
	if err != nil {
		logrus.Errorf("failed to listen: %v", err)
		return
	}

	db := repo.Conn()
	cache := repo.RedisClient()

	logrus.Infof("starting redis client: %v", cache.Ping(context.Background()).Val())

	defer cache.Close()

	s := grpc.NewServer()
	srv := &shorter.GRPCServer{DB: db, Cache: cache}
	proto.RegisterShorterServer(s, srv)

	logrus.Infof("Starting gRPC listener on port " + os.Getenv("GRPC_PORT"))
	if err := s.Serve(l); err != nil {
		logrus.Fatal(err)
	}
}
