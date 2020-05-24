package main

import (
	"github.com/coke00/challenge/controllers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", controllers.ShowWelcome).Methods("GET")
	r.HandleFunc("/api/contagiados", controllers.GetContagiados).Methods("GET")
	r.HandleFunc("/api/contagiados/{iso}", controllers.GetContagiadosPais).Methods("GET")

	log.Fatal(http.ListenAndServe(":4203", r))
}
