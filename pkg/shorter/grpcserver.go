package shorter

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/yherasymets/go-shorter/api/proto"
	"github.com/yherasymets/go-shorter/db"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GRPCServer...
type GRPCServer struct {
	proto.UnimplementedShorterServer
}

const (
	alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func shorting() string {
	b := make([]byte, 5)
	for i := range b {
		b[i] = alphabet[rand.Intn(len(alphabet))]
	}
	return string(b)
}

// TODO add case without http(s)

// func isValidUrl(token string) bool {
// 	_, err := url.ParseRequestURI(token)
// 	if err != nil {
// 		return false
// 	}
// 	u, err := url.Parse(token)
// 	if err != nil || u.Host == "" {
// 		return false
// 	}

// 	if u[:5] == ""

// 	return true
// }

func validateURL(u string) error {
	valid := govalidator.IsRequestURL(u)
	if !valid {
		return status.Error(codes.InvalidArgument, "invalid url")
	}
	return nil
}

func (GRPCServer) Short(ctx context.Context, req *proto.UrlRequest) (*proto.UrlResponse, error) {
	var link Link
	db := db.Conn()
	token := shorting()
	if err := validateURL(req.UrlName); err != nil {
		return nil, err
	}
	link.FullLink = req.UrlName
	link.Alias = token
	link.CreatedAt = time.Now()
	db.Create(&link)
	return &proto.UrlResponse{
		UrlResult: fmt.Sprintf("localhost:8081/%s", token),
	}, nil
}
