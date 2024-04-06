package main

import (
	"log"
	"net/http"
	"os"

	"github.com/yherasymets/go-shorter/internal/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	gRPCport = os.Getenv("GRPC_PORT")
	port     = os.Getenv("PORT")
)

func main() {
	clientGRPC, err := grpc.Dial(gRPCport, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect %v", err)
	}

	clientApp := api.NewApp(clientGRPC)
	handler := clientApp.Handler()

	log.Printf("starting client on port %v", port)
	if err = http.ListenAndServe(port, handler); err != nil {
		log.Fatalf("failed to connect %v: %v", port, err)
	}
}
