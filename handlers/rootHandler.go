package handlers

import "net/http"

func RootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not Found, To run url shortener user /url endpoint."))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Welcome to the URL Shortener API, To run url shortener user /url endpoint."))
}
