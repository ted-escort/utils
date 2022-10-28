package utils

import (
	"net/http"
)

func GetTLS(r *http.Request) string {
	if r.TLS == nil {
		return "http://"
	} else {
		return "https://"
	}
}