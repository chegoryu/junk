package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	config := LoadConfig()

	mux := http.NewServeMux()

	mux.HandleFunc("/headers", GetHeaders)

	addr := fmt.Sprintf(":%d", config.Port)
	log.Printf("starting server on %s", addr)

	err := http.ListenAndServe(addr, mux)
	if err != nil {
			log.Fatalf("Failed to start server: %v", err)
	}
}
