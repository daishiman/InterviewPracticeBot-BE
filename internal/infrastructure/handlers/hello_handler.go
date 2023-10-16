package handlers

import (
	"log"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello, world!"))
	if err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}
