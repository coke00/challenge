package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	// Route of Welcome
	r.HandleFunc("/", showWelcome).Methods("GET")
	r.HandleFunc("/api/contagiados", getContagiados).Methods("GET")
	r.HandleFunc("/api/contagiados/{iso}", getContagiadosPais).Methods("GET")

	log.Fatal(http.ListenAndServe(":4203", r))
}
