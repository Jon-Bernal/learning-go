package main

import (
	"log"
	"net/http"
)

func main() {
	// log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir("/Users/jon/Documents/web-dev/learning/golang/client"))))

	mux := http.NewServeMux()
	// Create sample handler to return 404 pages
	mux.Handle("/", mainPageHandler())
	mux.Handle("/api/*")

	// 404 Page
	mux.Handle("/*", http.NotFoundHandler())

	log.Fatal(http.ListenAndServe(":8080", mux))

}
