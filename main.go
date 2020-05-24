package main

import (
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", showWelcome).Methods("GET")
	r.HandleFunc("/api/contagiados", getContagiados).Methods("GET")
	r.HandleFunc("/api/contagiados/{iso}", getContagiadosPais).Methods("GET")

	log.Fatal(http.ListenAndServe(":4203", r))
}
func goDotEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}
