package shorter

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"github.com/yherasymets/go-shorter/pkg/utils"
	"github.com/yherasymets/go-shorter/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

const (
	charNumber     = 5
	statusExist    = "link already exist"
	statusSuccsess = "succsess"
)

type GRPCServer struct {
	proto.UnimplementedShorterServer
	DB    *gorm.DB
	Cache *redis.Client
}

func (g *GRPCServer) Create(ctx context.Context, req *proto.UrlRequest) (*proto.UrlResponse, error) {
	link := Link{}
	alias := utils.Shorting()
	if err := utils.ValidateURL(req.FullURL); err != nil {
		return nil, err
	}
	g.DB.Table("links").Where("full_link = ?", req.FullURL).Find(&link)
	if req.FullURL == link.FullLink {
		return &proto.UrlResponse{
			Result: fmt.Sprintf("localhost:8080/%s", link.Alias),
			Status: statusExist,
		}, nil
	}
	link.FullLink = req.FullURL
	link.Alias = alias
	link.CreatedAt = time.Now()
	g.DB.Create(&link)
	return &proto.UrlResponse{
		Result: fmt.Sprintf("localhost:8080/%s", link.Alias),
		Status: statusSuccsess,
	}, nil
}

func (g *GRPCServer) Get(ctx context.Context, req *proto.UrlRequest) (*proto.UrlResponse, error) {
	link := Link{}
	if req.FullURL == "" {
		return nil, status.Error(codes.InvalidArgument, "url must be set")
	}
	alias := req.FullURL[len(req.FullURL)-charNumber:]
	res, err := g.Cache.Get(ctx, alias).Result()
	if err == nil {
		link.FullLink = res
		logrus.Info("retrive from cache")
	} else {
		g.DB.Table("links").Where("alias = ?", alias).Find(&link)
		logrus.Info("retrive from DB")
		err := g.Cache.Set(ctx, alias, link.FullLink, time.Hour).Err()
		if err != nil {
			return nil, err
		}
	}
	return &proto.UrlResponse{
		Result: link.FullLink,
		Status: statusSuccsess,
	}, nil
}
