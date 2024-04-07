package utils

import (
	"errors"
	"math/rand"
	"regexp"

	"github.com/asaskevich/govalidator"
)

// Utils constants.
const (
	alphabet   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	pattern    = `^[a-zA-Z0-9]+$`
	charNumber = 5
)

// Shorting generates a random string of characters.
func Shorting() string {
	alias := make([]byte, charNumber)
	for i := range alias {
		alias[i] = alphabet[rand.Intn(len(alphabet))]
	}
	return string(alias)
}

// ValidateURL validates the given URL string.
func ValidateURL(url string) error {
	if url == "" {
		return errors.New("empty url")
	}
	if !govalidator.IsRequestURL(url) {
		return errors.New("invalid url")
	}
	return nil
}

// IsValidPath checks if the given path string matches the regular expression pattern.
func IsValidPath(path string) bool {
	// Compile the regular expression.
	re, err := regexp.Compile(pattern)
	// Return true if the compilation is successful and the path matches the pattern.
	return err == nil && re.MatchString(path)
}
