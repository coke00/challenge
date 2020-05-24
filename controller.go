package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)


func getContagiados(writer http.ResponseWriter, request *http.Request) {
	record, done := obtenerGlobalContagios()
	if done {
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	fmt.Printf("mostrando contagiados")
	crowd := record.Countries
	needle := "CL"

	respuesta := Find(crowd, needle)

	respuestaJson := &Contagiados{Data: &respuesta, Error: "Pais no encontrado"}
	fmt.Println("mostrando contagiados", respuestaJson)

	json.NewEncoder(writer).Encode(record)
}

func obtenerGlobalContagios() (retrieveSumary, bool) {
	setEnv()
	key := "UrlSumary"
	url, ex := os.LookupEnv(key)
	if !ex {
		log.Printf("La variable env %s no está establecida.\n", key)
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("NewRequest: ", err)
		return retrieveSumary{}, true
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("Do: ", err)
		return retrieveSumary{}, true
	}
	defer resp.Body.Close()
	var record retrieveSumary
	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}
	fmt.Println("Global  = ", record.Global)
	return record, false
}

func setEnv() {
	os.Setenv("UrlSumary", "https://api.covid19api.com/summary")
	fmt.Println("UrlSumary:", os.Getenv("UrlSumary"))
}


func Find(crowd []Countries, x string) Data {
	var std Data
	for i, v := range crowd {
		if v.CountryCode == x {
			fmt.Println("\n country  = ", v)
			fmt.Println("\n index  = ", i)
			std := Data{Country : v.Country, Iso : v.CountryCode, Confirmed : v.TotalConfirmed, NewConfirmed :v.NewConfirmed, Deaths : v.TotalDeaths}
			fmt.Println("std = ", std)
			return std
		}
	}
	fmt.Println("std-for-exit = ", std)
	return std
}
func showWelcome(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	fmt.Printf("hello")
	json.NewEncoder(writer).Encode("welcome")
}

func getContagiadosPais(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	fmt.Printf("mostrando contagiados por pais")
	params := mux.Vars(request)
	record, done := obtenerGlobalContagios()
	if done {
		return
	}
	crowd := record.Countries
	needle := params["iso"]
	respuesta := Find(crowd, needle)
	var contagiados Contagiados
	if respuesta.Iso != needle{
		contagiados.Error = "Pais no encontrado"
	}else{
		contagiados.Data = &respuesta
	}
	json.NewEncoder(writer).Encode(contagiados)
}
