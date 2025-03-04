package main

// CountryInfo represents general information about a country.
type CountryInfo struct {
	Name       string            `json:"name"`
	Continents []string          `json:"continents"`
	Population int               `json:"population"`
	Languages  map[string]string `json:"languages"`
	Borders    []string          `json:"borders"`
	Flag       string            `json:"flag"`
	Capital    string            `json:"capital"`
	Cities     []string          `json:"cities"`
}

// PopulationRecord represents a single year's population data.
type PopulationRecord struct {
	Year  int `json:"year"`
	Value int `json:"value"`
}

// PopulationResponse represents the response for the /population/ endpoint.
type PopulationResponse struct {
	Mean   int                `json:"mean"`
	Values []PopulationRecord `json:"values"`
}

// StatusResponse represents diagnostic information for the service.
type StatusResponse struct {
	CountriesNowAPI  interface{} `json:"countriesnowapi"`
	RestCountriesAPI interface{} `json:"restcountriesapi"`
	Version          string      `json:"version"`
	Uptime           int64       `json:"uptime"`
}
