package main

type Data struct {
	Country      string `json:"country"`
	Iso          string `json:"iso"`
	Confirmed    int    `json:"confirmed"`
	NewConfirmed int    `json:"newConfirmed"`
	Deaths       int    `json:"deaths"`
}

type Contagiados struct {
	Data  *Data  `json:"data"`
	Error string `json:"error"`
}

type retrieveSumary struct {
	Global    *Global     `json:"Global"`
	Countries []Countries `json:"Countries"`
	Date      string      `json:"Date"`
}
type Countries struct {
	Country        string `json:"Country"`
	CountryCode    string `json:"CountryCode"`
	Slug           string `json:"Slug"`
	NewConfirmed   int    `json:"NewConfirmed"`
	TotalConfirmed int    `json:"TotalConfirmed"`
	NewDeaths      int    `json:"NewDeaths"`
	TotalDeaths    int    `json:"TotalDeaths"`
	NewRecovered   int    `json:"NewRecovered"`
	TotalRecovered int    `json:"TotalRecovered"`
	Date           string `json:"Date"`
}
type Global struct {
	NewConfirmed   int `json:"NewConfirmed"`
	TotalConfirmed int `json:"TotalConfirmed"`
	NewDeaths      int `json:"NewDeaths"`
	TotalDeaths    int `json:"TotalDeaths"`
	NewRecovered   int `json:"NewRecovered"`
	TotalRecovered int `json:"TotalRecovered"`
}
