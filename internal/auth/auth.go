package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no auth found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("auth header format not correct")
	}
	if vals[0] != "ApiKey" {
		return "", errors.New("auth header first value format not correct")
	}

	return vals[1], nil
}