package main

import (
	"html/template"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Bog")
	ts, err := template.ParseFiles("/Users/kneehead/go_szef/letmecook/ui/public/index.html")
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}

func about(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("./ui/public/about.html")
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
