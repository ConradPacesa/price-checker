package handlers

import (
	"io"
	"net/http"
)

// Index is a check to see if the site is working
func Index(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
	}
	io.WriteString(w, "Status 200:ok")
}
