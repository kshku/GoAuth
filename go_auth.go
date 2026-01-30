package main

import (
	"net/http"
	"log"
	"os"

	"go_auth/internal/routes"
	"go_auth/internal/db"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	var port string = os.Getenv("PORT")
	if port == "" {
		log.Fatal("Failed to read PORT env var")
	}

	if err := db.InitDB(); err != nil {
		log.Fatalf("Failed to init db: %w", err)
	}

	mux := http.NewServeMux()

	routes.Register(mux)

	log.Println("starting server on :" + port)
	http.ListenAndServe(":" + port, mux)
}
