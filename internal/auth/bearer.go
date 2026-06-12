package auth

import (
	"errors"
	"net/http"
	"strings"
)

func GetBearerToken(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")

	if authHeader == "" {
		return "", errors.New("missing authorization header")
	}

	const prefix = "Bearer "

	if !strings.HasPrefix(authHeader, prefix) {
		return "", errors.New("invalid authorization header")
	}

	return strings.TrimPrefix(authHeader, prefix), nil
}
