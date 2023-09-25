package main

import (
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
	s := grpc.NewServer()
	srv := &shorter.GRPCServer{DB: db}
	proto.RegisterShorterServer(s, srv)

	logrus.Infof("Starting gRPC listener on port " + os.Getenv("GRPC_PORT"))
	if err := s.Serve(l); err != nil {
		logrus.Fatal(err)
	}
}
