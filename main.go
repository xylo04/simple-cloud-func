package main

import (
	"log"
	"net/http"
)

// Local development driver, not used by Cloud Functions
func main() {
	http.HandleFunc("/Hello", Hello)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func Hello(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Hello, world!"))
}
