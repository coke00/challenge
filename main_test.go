package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestObtenerGlobalContagios(t *testing.T) {
	want := "Hello, world."
	needle := "CL"
	assert := assert.New(t)
	var countrys []Countries
	countrys = append(countrys, Countries{Country: "Canada", CountryCode: "CA", Slug: "canada", NewConfirmed: 2410, TotalConfirmed: 167892, NewDeaths: 186, TotalDeaths: 12718, NewRecovered: 877, TotalRecovered: 42608, Date: "2020-05-24T02:02:54Z"})
	countrys = append(countrys, Countries{Country: "Chile", CountryCode: "CL", Slug: "cape-verde", NewConfirmed: 6, TotalConfirmed: 362, NewDeaths: 0, TotalDeaths: 4, NewRecovered: 0, TotalRecovered: 95, Date: "2020-05-24T02:02:54Z"})

	if respuesta := Find(countrys, needle); respuesta.Iso != needle {
		t.Errorf("Hello() = %q, want %q", respuesta, want)
	}else{

	assert.Equal(respuesta.Iso, needle, "el pais contiene informacion")
	}
}
