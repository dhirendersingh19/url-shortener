package main

import (
	"fmt"
	"net/http"
	"os"
	"url-shortener/handlers"
)

func main() {
	http.HandleFunc("/", handlers.RootHandler)
	http.HandleFunc("/url", handlers.UrlRouter)
	err := http.ListenAndServe("127.0.0.1:8081", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
