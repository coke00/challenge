package main

import (
	"github.com/coke00/challenge/controllers"
	"github.com/coke00/challenge/handlers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestObtenerGlobalContagios(t *testing.T) {
	needle := "CL"
	assert := assert.New(t)
	var countrys []handlers.Countries
	countrys = append(countrys, handlers.Countries{Country: "Canada", CountryCode: "CA", Slug: "canada", NewConfirmed: 2410, TotalConfirmed: 167892, NewDeaths: 186, TotalDeaths: 12718, NewRecovered: 877, TotalRecovered: 42608, Date: "2020-05-24T02:02:54Z"})
	countrys = append(countrys, handlers.Countries{Country: "Chile", CountryCode: "CL", Slug: "cape-verde", NewConfirmed: 6, TotalConfirmed: 362, NewDeaths: 0, TotalDeaths: 4, NewRecovered: 0, TotalRecovered: 95, Date: "2020-05-24T02:02:54Z"})

	if respuesta := controllers.Find(countrys, needle); respuesta.Iso != needle {
		t.Errorf("TestObtenerGlobalContagios() = %q, needle %q", respuesta.Iso, needle)
	}else{

	assert.Equal(respuesta.Iso, needle, "el pais contiene informacion")
	}
}
