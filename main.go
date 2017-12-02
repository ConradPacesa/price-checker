package main

import (
	"net/http"

	"github.com/ConradPacesa/price-checker/handlers"
)

func main() {
	http.HandleFunc("/", handlers.Index)
	http.HandleFunc("/search", handlers.Search)
	http.ListenAndServe(":80", nil)
}
