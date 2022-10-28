package utils

import (
	"net/http"
)

func TLS() string {
	var r http.Request
	if r.TLS == nil {
		return "http://"
	} else {
		return "https://"
	}
}
