package main

import (
	"html/template"
	"log"
	"net/http"
)

func about(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("./ui/pages/about.html")
	if err != nil {
		log.Print(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = ts.Execute(w, ts)
	if err != nil {
		log.Print(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	ts, err := template.ParseFiles("./ui/pages/contact.html")
	if err != nil {
		log.Print(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = ts.Execute(w, ts)
	if err != nil {
		log.Print(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
