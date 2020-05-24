package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/coke00/challenge/handlers"
	"github.com/coke00/challenge/utils"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)


func GetContagiados(writer http.ResponseWriter, request *http.Request) {
	record, done, error := obtenerContagiosGlobales()
	if done {
		var contagiados handlers.Contagiados
		contagiados.Error = error.Error()
		json.NewEncoder(writer).Encode(contagiados)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	fmt.Printf("mostrando contagiados")
	json.NewEncoder(writer).Encode(record)
}

func obtenerContagiosGlobales() (handlers.RetrieveSumary, bool, error) {
	url := utils.GoDotEnvVariable("UrlSumary")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println("nuevoRequest: ", err)
		return handlers.RetrieveSumary{}, true, err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Obteniendo: ", err)
		return handlers.RetrieveSumary{}, true, err
	}
	defer resp.Body.Close()
	var record handlers.RetrieveSumary
	if err := json.NewDecoder(resp.Body).Decode(&record); err != nil {
		log.Println(err)
	}
	fmt.Println("Global  = ", record.Global)
	return record, false, nil
}

func Find(crowd []handlers.Countries, x string) handlers.Data {
	var std handlers.Data
	for i, v := range crowd {
		if v.CountryCode == x {
			fmt.Println("index  = ", i)
			std := handlers.Data{Country : v.Country, Iso : v.CountryCode, Confirmed : v.TotalConfirmed, NewConfirmed :v.NewConfirmed, Deaths : v.TotalDeaths}
			fmt.Println("std = ", std)
			return std
		}
	}
	fmt.Println("std-for-exit = ", std)
	return std
}
func ShowWelcome(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	fmt.Printf("hello")
	json.NewEncoder(writer).Encode("welcome")
}

func GetContagiadosPais(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	fmt.Println("mostrando contagiados por pais")
	params := mux.Vars(request)
	var contagiados handlers.Contagiados
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
		contagiados.Error = "codigo de pais no existe"
	}else{
		contagiados.Data = &respuesta
	}
	json.NewEncoder(writer).Encode(contagiados)
}
