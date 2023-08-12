package main

import (
	"log"
	"net"

	"github.com/yherasymets/go-shorter/api/proto"
	"github.com/yherasymets/go-shorter/pkg/shorter"
	"google.golang.org/grpc"
)

const port = ":8081"

func main() {
	l, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	srv := &shorter.GRPCServer{}
	proto.RegisterShorterServer(s, srv)

	log.Printf("Starting gRPC listener on port " + port)

	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
