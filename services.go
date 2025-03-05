package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

// restCountriesAPI is the base URL for the REST Countries API.
const restCountriesAPI = "http://129.241.150.113:8080/v3.1/"

// restCountryResponse is a struct that matches the JSON structure returned by the REST Countries API.
type restCountryResponse struct {
	Name struct {
		Common   string `json:"common"`
		Official string `json:"official"`
	} `json:"name"`
	Continents []string          `json:"continents"`
	Population int               `json:"population"`
	Languages  map[string]string `json:"languages"`
	Borders    []string          `json:"borders"`
	Capital    []string          `json:"capital"`
	Flags      struct {
		Png string `json:"png"`
		Svg string `json:"svg"`
	} `json:"flags"`
}

// fetchCountryInfoFromRest fetches country info from the REST Countries API based on a two-letter country code.
// It converts the API response into our CountryInfo struct.
func fetchCountryInfoFromRest(countryCode string) (*CountryInfo, error) {
	url := restCountriesAPI + "alpha/" + countryCode
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to fetch data from REST Countries API")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var restData []restCountryResponse
	if err := json.Unmarshal(body, &restData); err != nil {
		return nil, err
	}
	if len(restData) == 0 {
		return nil, errors.New("no country data found")
	}

	rc := restData[0]
	cInfo := &CountryInfo{
		Name:       rc.Name.Common,
		Continents: rc.Continents,
		Population: rc.Population,
		Languages:  rc.Languages,
		Borders:    rc.Borders,
		Flag:       rc.Flags.Png,
	}
	if len(rc.Capital) > 0 {
		cInfo.Capital = rc.Capital[0]
	} else {
		cInfo.Capital = ""
	}

	return cInfo, nil
}

// countriesNowAPI is the base URL for the CountriesNow API.
const countriesNowAPI = "http://129.241.150.113:3500/api/v0.1/"

// CountriesNowCitiesResponse defines the JSON structure returned by the CountriesNow API for cities.
type CountriesNowCitiesResponse struct {
	Error bool     `json:"error"`
	Msg   string   `json:"msg"`
	Data  []string `json:"data"`
}

// fetchCitiesFromCountriesNow fetches the list of cities for a given country name using the CountriesNow API.
// The API expects a POST request with a JSON body containing "country": <countryName>.
func fetchCitiesFromCountriesNow(countryName string) ([]string, error) {
	url := countriesNowAPI + "countries/cities"
	payload, err := json.Marshal(map[string]string{
		"country": countryName,
	})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to fetch cities from CountriesNow API")
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result CountriesNowCitiesResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	if result.Error {
		return nil, errors.New("CountriesNow API error: " + result.Msg)
	}
	return result.Data, nil
}

// populationDataMap simulates an in-memory database of historical population data.
// NOTE: Currently, it only has data for "NO" (Norway). This can be extended
// or replaced with a real external data source if needed.

var populationDataMap = map[string][]PopulationRecord{
	"NO": {
		{Year: 2010, Value: 4889252},
		{Year: 2011, Value: 4953088},
		{Year: 2012, Value: 5018573},
		{Year: 2013, Value: 5079623},
		{Year: 2014, Value: 5137232},
		{Year: 2015, Value: 5188607},
	},
	// Additional countries can be added here.
}

// getPopulationHistory returns historical population data for a given country code from the in-memory map.
func getPopulationHistory(countryCode string) ([]PopulationRecord, error) {
	if data, ok := populationDataMap[countryCode]; ok {
		return data, nil
	}
	return nil, fmt.Errorf("no population data found for country code %s", countryCode)
}

// checkRestCountriesAPI performs a GET request to check if the REST Countries API is reachable.
// It returns the HTTP status code, or 0 if not reachable.
func checkRestCountriesAPI() int {
	url := restCountriesAPI + "alpha/NO"
	resp, err := http.Get(url)
	if err != nil {
		return 0
	}
	defer resp.Body.Close()
	io.Copy(io.Discard, resp.Body)
	return resp.StatusCode
}

// checkCountriesNowAPI performs a POST request to check if the CountriesNow API is reachable.
// It returns the HTTP status code, or 0 if not reachable.

func checkCountriesNowAPI() int {
	url := countriesNowAPI + "countries/cities"
	payload, err := json.Marshal(map[string]string{
		"country": "Norway",
	})
	if err != nil {
		return 0
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(payload))
	if err != nil {
		return 0
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0
	}
	defer resp.Body.Close()
	io.Copy(io.Discard, resp.Body)
	return resp.StatusCode
}
