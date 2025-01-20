package server

import (
	"net/http"

	"github.com/nepalagroyantra/ovogenix/api/egg"
)

// RegisterRoutes sets up the application's routes
func RegisterRoutes(mux *http.ServeMux) {
    mux.HandleFunc("/egg", egg.Handler)
}