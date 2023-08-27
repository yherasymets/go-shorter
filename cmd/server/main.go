package main

import (
	"log"
	"net"
	"os"

	"github.com/yherasymets/go-shorter/databases"
	"github.com/yherasymets/go-shorter/pkg/shorter"
	"github.com/yherasymets/go-shorter/proto"
	"google.golang.org/grpc"
)

func main() {
	l, err := net.Listen("tcp", os.Getenv("GRPC_PORT"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	db := databases.Conn()
	s := grpc.NewServer()
	srv := &shorter.GRPCServer{
		DB: db,
	}
	proto.RegisterShorterServer(s, srv)
	log.Printf("Starting gRPC listener on port " + os.Getenv("GRPC_PORT"))
	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
