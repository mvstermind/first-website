package main

import (
	"log"
	"net/http"
)

func meth() {

	mux := http.NewServeMux()

	mux.HandleFunc("/{$}", home)

	log.Print("Starting server on :2121")

	err := http.ListenAndServe(":2121", mux)
	if err != nil {
		log.Fatal(err)
		return
	}

}

func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./web/pages/index.html")
	fs := http.FileServer(http.Dir("./web/styles"))
	mux.Handle("/styles/", http.StripPrefix("/styles", fs))
}
