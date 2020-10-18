package gcf

import "net/http"

// Cloud function
func Hello(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Hello, world!"))
}
