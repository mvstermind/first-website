package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/pages/"))
	mux.Handle("GET /pages/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("GET /", home)

	log.Print("SERVER YAHOO ON :2137")

	err := http.ListenAndServe(":2137", mux)
	if err != nil {
		log.Print("something went wrong homie")
		return
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./ui/pages/index.html")
}

