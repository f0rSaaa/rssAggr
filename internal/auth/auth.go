package auth

import (
	"errors"
	"net/http"
	"strings"
)

//header Authorization = ApiKey {actual api key}
func GetApiKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")

	if val == "" {
		return "", errors.New("no authenticatoin info found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("malformed headers")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("malformed first part of the header")
	}

	return vals[1], nil
}
