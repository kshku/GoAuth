package routes

import (
	"net/http"

	"go_auth/internal/handlers"
	"go_auth/internal/middlewares"
)

func applyMiddlewares(handler http.HandlerFunc, middlewares ...func(http.HandlerFunc) http.HandlerFunc) http.HandlerFunc {
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}
	return handler
}

func Register(mux *http.ServeMux) {
	mux.HandleFunc("/health", handlers.Health)

	commonMiddlewares := []func(http.HandlerFunc) http.HandlerFunc{
		middlewares.LoggingMiddleware,
		middlewares.CORSMiddleware,
	}

	publicMiddlewares := append(commonMiddlewares, middlewares.RateLimitMiddleware)
	// protectedMiddlewares := append(commonMiddlewares, middlewares.AuthMiddleware)

	mux.HandleFunc("/register", applyMiddlewares(handlers.Register, publicMiddlewares...))
	mux.HandleFunc("/login", applyMiddlewares(handlers.Login, publicMiddlewares...))
}
