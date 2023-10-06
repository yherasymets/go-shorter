package main

import (
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/yherasymets/go-shorter/pkg/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		logrus.Fatalf("failed to connect :8081 %v", err)
	}

	client := api.ClientApp{
		Conn: conn,
	}

	logrus.Info("starting client..")
	if err = http.ListenAndServe(":8080", client.Routes()); err != nil {
		logrus.Fatalf("failed to connect :8080 %v", err)
	}

}
