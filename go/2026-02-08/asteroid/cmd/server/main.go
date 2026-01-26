package main

import (
	"log"

	"asteroid/internal/http"
)

func main() {
	server := http.NewServer()

	log.Println("Starting server on :8080")
	if err := server.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
