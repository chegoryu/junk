package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/chegoryu/junk/go/cmd/echo_server/config"
	"github.com/chegoryu/junk/go/cmd/echo_server/handlers"
)

func main() {
	cfg := config.LoadConfig()

	mux := http.NewServeMux()
	handlers.AddHandlers(mux)

	addr := fmt.Sprintf(":%d", cfg.Port)
	log.Printf("starting server on %s", addr)

	err := http.ListenAndServe(addr, mux)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
