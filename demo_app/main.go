package main

import (
	"log"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(PlayerServer)
	if err := http.ListenAndServe(":8000", handler); err != nil {
		log.Fatalf("could not listen on port 8000 %v", err)
	}
}
