package main

import (
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/yherasymets/go-shorter/internal/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	gRPCport = os.Getenv("GRPC_PORT")
	port     = os.Getenv("PORT")
)

func main() {
	conn, err := grpc.Dial(gRPCport, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logrus.Fatalf("failed to connect %v", err)
	}

	client := api.NewApp(conn)
	handler := client.Handler()

	logrus.Info("starting client on port ", port)
	if err = http.ListenAndServe(port, handler); err != nil {
		logrus.Fatalf("failed to connect %v: %v", port, err)
	}
}
