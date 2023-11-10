package utils

import (
	"math/rand"

	"github.com/asaskevich/govalidator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	alphabet   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	charNumber = 5
)

func Shorting() string {
	alias := make([]byte, charNumber)
	for i := range alias {
		alias[i] = alphabet[rand.Intn(len(alphabet))]
	}
	return string(alias)
}

func ValidateURL(url string) error {
	if url == "" {
		return status.Error(codes.InvalidArgument, "empty url")
	}
	if !govalidator.IsRequestURL(url) {
		return status.Error(codes.InvalidArgument, "invalid url")
	}
	return nil
}
