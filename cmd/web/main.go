package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/{$}", home)
	mux.HandleFunc("/about", about)

	log.Print("Starting new server on port :2137")

	err := http.ListenAndServe(":2137", mux)
	log.Fatal(err)
}
