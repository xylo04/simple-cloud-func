package main

import (
	gcf "github.com/xylo04/simple-cloud-func"
	"log"
	"net/http"
)

// Local development driver, not used by Cloud Functions
func main() {
	http.HandleFunc("/Hello", gcf.Hello)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
