package server

import (
	"fmt"
	"log"
	"net/http"
)

// StartServer initializes and starts the HTTP server
func StartServer() {
    mux := http.NewServeMux()
    RegisterRoutes(mux)

    port := ":8003"
    fmt.Println("Starting server on port", port)
    err := http.ListenAndServe(port, mux)
    if err != nil {
        log.Fatalf("Server failed to start: %v", err)
    }
}
