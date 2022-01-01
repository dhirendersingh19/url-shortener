package handlers

import (
	"net/http"
	"strings"
)

func UrlRouter(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimSuffix(r.URL.Path, "/")
	if path == "/url" {
		switch r.Method {
		case http.MethodGet:
			return
		case http.MethodPost:
			return
		default:
			postError(w, http.StatusMethodNotAllowed)
			return
		}
	}
}
