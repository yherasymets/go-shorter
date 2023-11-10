package utils

import (
	"math/rand"
	"regexp"

	"github.com/asaskevich/govalidator"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	alphabet   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	pattern    = `^[a-zA-Z0-9]+$`
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

func IsValidPath(path string) bool {
	// Compile the regular expression.
	re, err := regexp.Compile(pattern)
	// Return true if the compilation is successful and the path matches the pattern.
	return err == nil && re.MatchString(path)
}
