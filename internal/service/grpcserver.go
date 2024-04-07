package service

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/yherasymets/go-shorter/pkg/utils"
	"github.com/yherasymets/go-shorter/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

const (
	charNumber     = 5
	statusExist    = "link already exist"
	statusSuccsess = "operation succsess"
)

type GRPCServer struct {
	proto.UnimplementedShorterServer
	DB    *gorm.DB
	Cache *redis.Client
}

// type Service interface {
// 	Create(ctx context.Context, req *proto.CreateRequest) (*proto.UrlResponse, error)
// 	Get(ctx context.Context, req *proto.GetRequest) (*proto.UrlResponse, error)
// }

func NewService(db *gorm.DB, cache *redis.Client) *GRPCServer {
	return &GRPCServer{
		UnimplementedShorterServer: proto.UnimplementedShorterServer{},
		DB:                         db,
		Cache:                      cache,
	}
}

func (s *GRPCServer) Create(ctx context.Context, req *proto.CreateRequest) (*proto.UrlResponse, error) {
	var link Link
	alias := utils.Shorting()
	if err := utils.ValidateURL(req.FullURL); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if err := s.DB.Table("links").Where("full_link = ?", req.FullURL).Find(&link).Error; err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	if req.FullURL == link.FullLink {
		return &proto.UrlResponse{
			Result: fmt.Sprintf("localhost:8080/%s", link.Alias),
			Status: statusExist,
		}, nil
	}

	link.FullLink = req.FullURL
	link.Alias = alias
	link.CreatedAt = time.Now()
	if err := s.DB.Create(&link).Error; err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &proto.UrlResponse{
		Result: fmt.Sprintf("localhost:8080/%s", link.Alias),
		Status: statusSuccsess,
	}, nil
}

func (s *GRPCServer) Get(ctx context.Context, req *proto.GetRequest) (*proto.UrlResponse, error) {
	var link Link
	if req.ShortURL == "" {
		return nil, status.Error(codes.InvalidArgument, "url must be set")
	}
	alias := req.ShortURL[len(req.ShortURL)-charNumber:]
	res, err := s.Cache.Get(ctx, alias).Result()
	if err == nil {
		link.FullLink = res
		log.Println("retrive from cache", res)
	} else {
		if err := s.DB.Table("links").Where("alias = ?", alias).Find(&link).Error; err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
		log.Printf("retrive from DB %+v", link)
		err := s.Cache.Set(ctx, alias, link.FullLink, time.Hour).Err()
		if err != nil {
			return nil, status.Error(codes.Internal, err.Error())
		}
	}
	return &proto.UrlResponse{
		Result: link.FullLink,
		Status: statusSuccsess,
	}, nil
}
