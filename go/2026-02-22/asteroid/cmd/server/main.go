package main

import (
	"log"

	"asteroid/internal/http"
)

func main() {
	issuer := "http://localhost:8080"
	server := http.NewServer(issuer)

	log.Println("Starting server on :8080")
	if err := server.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
