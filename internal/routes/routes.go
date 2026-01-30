package routes

import (
	"net/http"

	"go_auth/internal/handlers"
)

func Register(mux *http.ServeMux) {
	mux.HandleFunc("/health", handlers.Health)
}
