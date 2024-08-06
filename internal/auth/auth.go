package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey extracts the API key from the headers
// Format - Auth-Key: API_KEY <key>
func GetAPIKey(headers http.Header) (string, error) {
    val := headers.Get("Auth-Key")
    if val == "" {
        return "", errors.New("authentication header is missing")
    }

    authHeaders := strings.Split(val, " ")
    if len(authHeaders) != 2 {
        return "", errors.New("invalid authentication header")
    }

    if authHeaders[0] != "API_KEY" {
        return "", errors.New("invalid start to authentication header")
    }

    return authHeaders[1], nil
}