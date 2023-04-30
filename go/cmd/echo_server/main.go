package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/headers", GetHeaders)

	log.Fatal(http.ListenAndServe(":12365", mux))
}
