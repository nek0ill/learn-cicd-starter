package auth

import (
	"errors"
	"net/http"
	"strings"
	"testing"
)

var ErrNoAuthHeaderIncluded = errors.New("no authorization header included")

func GetAPIKey(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", ErrNoAuthHeaderIncluded
	}
	splitAuth := strings.Split(authHeader, " ")
	if len(splitAuth) < 2 || splitAuth[0] != "ApiKey" {
		return "", errors.New("malformed La-Li-Lu-Le-Lo header")
	}

	return splitAuth[1], nil
}

func TestSplit(t *testing.T) {
	t.Run("returns error when authorization header is missing", func(t *testing.T) {
		headers := http.Header{}
		_, err := GetAPIKey(headers)
		if err == nil {
			t.Fatalf("expected error, got nil")
		}
		if err != ErrNoAuthHeaderIncluded {
			t.Fatalf("expected error %v, got %v", ErrNoAuthHeaderIncluded, err)
		}
	})

	t.Run("returns error when authorization header is malformed", func(t *testing.T) {
		headers := http.Header{}
		headers.Set("Authorization", "Bearer my-api-key")
		_, err := GetAPIKey(headers)
		if err == nil {
			t.Fatalf("expected error, got nil")
		}
		if err.Error() != "malformed authorization header" {
			t.Fatalf("expected error 'malformed authorization header', got %v", err)
		}
	})
	t.Run("returns error when authorization header is empty", func(t *testing.T) {
		headers := http.Header{}
		headers.Set("Authorization", "")
		_, err := GetAPIKey(headers)
		if err == nil {
			t.Fatalf("expected error, got nil")
		}
		if err != ErrNoAuthHeaderIncluded {
			t.Fatalf("expected error %v, got %v", ErrNoAuthHeaderIncluded, err)
		}
	})
}
