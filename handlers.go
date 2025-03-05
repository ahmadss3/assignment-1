package main

import (
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"
)

// handleCountryInfo handles requests for the Country Info endpoint.
// It extracts the two-letter country code, fetches general country info
// from the REST Countries API, fetches city data from the CountriesNow API,
// and returns the combined data. It also checks if the user wants HTML output.

func handleCountryInfo(w http.ResponseWriter, r *http.Request) {
	countryCode := strings.TrimPrefix(r.URL.Path, "/countryinfo/v1/info/")
	if countryCode == "" {
		http.Error(w, "Country code is required", http.StatusBadRequest)
		return
	}
	countryCode = strings.ToUpper(countryCode)

	// Fetch general country info from REST Countries API.
	countryInfo, err := fetchCountryInfoFromRest(countryCode)
	if err != nil {
		http.Error(w, "Failed to fetch country info: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Fetch cities using the country's common name.
	cities, err := fetchCitiesFromCountriesNow(countryInfo.Name)
	if err != nil {
		http.Error(w, "Failed to fetch cities: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Sort cities alphabetically.
	sort.Strings(cities)

	// Optional query parameter to limit the number of cities returned.
	limitStr := r.URL.Query().Get("limit")
	if limitStr != "" {
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			http.Error(w, "Invalid limit parameter", http.StatusBadRequest)
			return
		}
		if limit > 0 && limit < len(cities) {
			cities = cities[:limit]
		}
	}

	// Assign the sorted (and possibly truncated) list of cities.
	countryInfo.Cities = cities

	// Respond in JSON or HTML, depending on the 'format' query parameter.
	if err := respondJSONOrHTML(w, r, countryInfo); err != nil {
		http.Error(w, "Failed to encode response: "+err.Error(), http.StatusInternalServerError)
	}
}

// handleCountryPopulation handles requests for the Country Population endpoint.
// It retrieves population data from an in-memory map, optionally filters it
// by a year range, calculates the mean, and returns the result.
// It also checks if the user wants HTML output.

func handleCountryPopulation(w http.ResponseWriter, r *http.Request) {
	countryCode := strings.TrimPrefix(r.URL.Path, "/countryinfo/v1/population/")
	if countryCode == "" {
		http.Error(w, "Country code is required", http.StatusBadRequest)
		return
	}
	countryCode = strings.ToUpper(countryCode)

	records, err := getPopulationHistory(countryCode)
	if err != nil {
		http.Error(w, "Error: "+err.Error(), http.StatusNotFound)
		return
	}

	// Check if the user specified a year range (?limit=startYear-endYear).
	limitStr := r.URL.Query().Get("limit")
	filteredRecords := records
	if limitStr != "" {
		startYear, endYear, err := parseYearRange(limitStr)
		if err != nil {
			http.Error(w, "Invalid limit parameter: "+err.Error(), http.StatusBadRequest)
			return
		}
		filteredRecords = []PopulationRecord{}
		for _, rec := range records {
			if rec.Year >= startYear && rec.Year <= endYear {
				filteredRecords = append(filteredRecords, rec)
			}
		}
		if len(filteredRecords) == 0 {
			http.Error(w, "No population data found for the given year range", http.StatusNotFound)
			return
		}
	}

	// Calculate the mean population for the filtered records.
	var total int
	for _, rec := range filteredRecords {
		total += rec.Value
	}
	mean := total / len(filteredRecords)

	response := PopulationResponse{
		Mean:   mean,
		Values: filteredRecords,
	}

	// Respond in JSON or HTML, depending on the 'format' query parameter.
	if err := respondJSONOrHTML(w, r, response); err != nil {
		http.Error(w, "Failed to encode response: "+err.Error(), http.StatusInternalServerError)
	}
}

// handleStatus handles requests for the Diagnostics/Status endpoint.
// It checks the status of both external APIs (REST Countries API and CountriesNow API)
// and calculates the service uptime. It also checks if the user wants HTML output.

func handleStatus(w http.ResponseWriter, r *http.Request) {
	restStatus := checkRestCountriesAPI()
	countriesNowStatus := checkCountriesNowAPI()

	uptime := int64(time.Since(startTime).Seconds())

	status := StatusResponse{
		CountriesNowAPI:  countriesNowStatus,
		RestCountriesAPI: restStatus,
		Version:          "v1",
		Uptime:           uptime,
	}

	// Respond in JSON or HTML, depending on the 'format' query parameter.

	if err := respondJSONOrHTML(w, r, status); err != nil {
		http.Error(w, "Failed to encode status response: "+err.Error(), http.StatusInternalServerError)
	}
}
