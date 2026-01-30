package main

import (
	"net/http"
	"log"
	"os"

	"go_auth/internal/routes"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	var port string = os.Getenv("PORT")
	if port == "" {
		panic("Port is not set")
	}

	mux := http.NewServeMux()

	routes.Register(mux)

	log.Println("starting server on :" + port)
	http.ListenAndServe(":" + port, mux)
}
