package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load("env")
	portstring := os.Getenv("PORT")
	router := chi.NewRouter()
	if portstring == "" {
		log.Fatal("port is not available in this system")
	}
	fmt.Printf(portstring)
}
