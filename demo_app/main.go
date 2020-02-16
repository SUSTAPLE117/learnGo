package main

import (
	"log"
	"net/http"
)

func main() {
	server := &PlayerServer{}
	if err := http.ListenAndServe(":8000", server); err != nil {
		log.Fatalf("could not listen on port 8000 %v", err)
	}
}
