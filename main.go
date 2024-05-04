package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	mux.HandleFunc("/", home)

	log.Print("Starting server on :4040")

	err := http.ListenAndServe(":4040", mux)
	if err != nil {
		log.Print("Error: ", err)
		return
	}

}

func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}
