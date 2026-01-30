package routes

import (
	"net/http"

	"go_auth/internal/handlers"
)

func Register(mux *http.ServeMux) {
	mux.HandleFunc("/health", handlers.Health)
	mux.HandleFunc("/register", handlers.Register)
	mux.HandleFunc("/login", handlers.Login)
}
