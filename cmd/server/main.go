package main

import (
	"InterviewPracticeBot-BE/internal/infrastructure/handlers"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.HelloHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
