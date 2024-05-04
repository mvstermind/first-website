package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("./ui"))
	mux.Handle("/static/", http.StripPrefix("/ui", fs))

	mux.HandleFunc("/{$}", home)
	mux.HandleFunc("/about", about)

	log.Print("Starting server on :4040")

	err := http.ListenAndServe(":4040", mux)
	if err != nil {
		log.Print("Error: ", err)
		return
	}

}

func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./ui/index.html")
}
func about(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./ui/about.html")
}
