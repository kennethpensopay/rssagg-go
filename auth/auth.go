package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey extracts an API Key from
// the headers of an HTTP request
// Example:
// Authorization: ApiKey {insert apikey here}
func GetAPIKey(h http.Header) (string, error) {
	if authValue := h.Get("Authorization"); authValue != "" {
		authParts := strings.Split(authValue, " ")

		if len(authParts) != 2 {
			return "", errors.New("malformed authentication header")
		} else if len(authParts) == 2 && authParts[0] != "ApiKey" {
			return "", errors.New("malformed first part of authentication header")
		}

		return authParts[1], nil
	}
	return "", errors.New("no authentication info found")
}
