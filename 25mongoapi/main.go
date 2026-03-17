package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/sh4shv4t/mongoapi/controller"
	"github.com/sh4shv4t/mongoapi/router"
)

func main() {
	fmt.Println("MongoDB API")
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, reading environment variables directly")
	}

	controller.ConnectDB()

	r := mux.NewRouter()
	router.Router(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	log.Printf("Server running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
