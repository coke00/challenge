package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)


func getContagiados(writer http.ResponseWriter, request *http.Request) {
	record, done, error := obtenerContagiosGlobales()
	if done {
		var contagiados Contagiados
		contagiados.Error = error.Error()
		json.NewEncoder(writer).Encode(contagiados)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	fmt.Printf("mostrando contagiados")
	json.NewEncoder(writer).Encode(record)
}

func obtenerContagiosGlobales() (retrieveSumary, bool, error) {
	url := goDotEnvVariable("UrlSumary")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("nuevoRequest: ", err)
		return retrieveSumary{}, true, err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Obteniendo: ", err)
		return retrieveSumary{}, true, err
	}
	defer resp.Body.Close()
	var record retrieveSumary
	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}
	fmt.Println("Global  = ", record.Global)
	return record, false, nil
}

func Find(crowd []Countries, x string) Data {
	var std Data
	for i, v := range crowd {
		if v.CountryCode == x {
			fmt.Println("index  = ", i)
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
	var contagiados Contagiados
	record, done, error := obtenerContagiosGlobales()
	if done {
		contagiados.Error = error.Error()
		json.NewEncoder(writer).Encode(contagiados)
		return
	}
	crowd := record.Countries
	needle := params["iso"]
	respuesta := Find(crowd, needle)
	if respuesta.Iso != needle{
		contagiados.Error = "Pais no encontrado"
	}else{
		contagiados.Data = &respuesta
	}
	json.NewEncoder(writer).Encode(contagiados)
}
