package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static", http.StripPrefix("/static", fs))

	mux.HandleFunc("/{$}", home)

	log.Print("SHOOOOO :2137")

	err := http.ListenAndServe(":2137", mux)
	if err != nil {
		log.Fatal(err)
		return
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	log.Print("Home page visited")
}
