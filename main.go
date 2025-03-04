package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// startTime holds the time when the server started (used for uptime calculation).
var startTime time.Time

func main() {
	// Record the start time.
	startTime = time.Now()

	fmt.Println("Starting server on port 8080...")

	// Register endpoints.
	http.HandleFunc("/countryinfo/v1/info/", handleCountryInfo)
	http.HandleFunc("/countryinfo/v1/population/", handleCountryPopulation)
	http.HandleFunc("/countryinfo/v1/status/", handleStatus)

	// Root endpoint for a welcome message.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to the Country Information Service!")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
