package shorter

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/yherasymets/go-shorter/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

const (
	alphabet   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	charNumber = 5

	statusExist    = "link already exist"
	statusSuccsess = "succsess"
)

type GRPCServer struct {
	proto.UnimplementedShorterServer
	DB *gorm.DB
}

func (g *GRPCServer) Create(ctx context.Context, req *proto.UrlRequest) (*proto.UrlResponse, error) {
	link := Link{}
	alias := shorting()
	if err := validateURL(req.FullURL); err != nil {
		return nil, err
	}

	g.DB.Table("links").Where("full_link = ?", req.FullURL).Find(&link)
	if req.FullURL == link.FullLink {
		return &proto.UrlResponse{
			Result: fmt.Sprintf("localhost:8081/%s", link.Alias),
			Status: statusExist,
		}, nil
	}

	link.FullLink = req.FullURL
	link.Alias = alias
	link.CreatedAt = time.Now()
	g.DB.Create(&link)

	return &proto.UrlResponse{
		Result: fmt.Sprintf("localhost:8081/%s", link.Alias),
		Status: statusSuccsess,
	}, nil
}

func (g *GRPCServer) Get(ctx context.Context, req *proto.UrlRequest) (*proto.UrlResponse, error) {
	var link Link

	if req.FullURL == "" {
		return nil, status.Error(codes.InvalidArgument, "url must be set")
	}

	g.DB.Table("links").Where("alias = ?", req.FullURL[len(req.FullURL)-charNumber:]).Find(&link)
	return &proto.UrlResponse{
		Result: link.FullLink,
		Status: statusSuccsess,
	}, nil
}

func shorting() string {
	alias := make([]byte, charNumber)
	for i := range alias {
		alias[i] = alphabet[rand.Intn(len(alphabet))]
	}
	return string(alias)
}

func validateURL(url string) error {
	if url == "" {
		return status.Error(codes.InvalidArgument, "empty url")
	}
	if url[:4] == "http" || url[:5] == "https" {
		valid := govalidator.IsRequestURL(url)
		if !valid {
			return status.Error(codes.InvalidArgument, "invalid url")
		}
	}
	return nil
}
