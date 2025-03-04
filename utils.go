package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

// parseYearRange parses a limit string in the format "startYear-endYear" (e.g., "2010-2015")
// and returns the corresponding start and end years as integers.
func parseYearRange(limitStr string) (int, int, error) {
	parts := strings.Split(limitStr, "-")
	if len(parts) != 2 {
		return 0, 0, fmt.Errorf("invalid limit format, expected startYear-endYear")
	}
	startYear, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, 0, fmt.Errorf("invalid start year: %v", err)
	}
	endYear, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, 0, fmt.Errorf("invalid end year: %v", err)
	}
	if startYear > endYear {
		return 0, 0, fmt.Errorf("start year must be less than or equal to end year")
	}
	return startYear, endYear, nil
}

// respondJSONOrHTML checks if the user requested ?format=html.
// If so, it returns an HTML page with the JSON data inside a <pre> block.
// Otherwise, it returns raw JSON with Content-Type: application/json.
func respondJSONOrHTML(w http.ResponseWriter, r *http.Request, data interface{}) error {
	format := r.URL.Query().Get("format")
	if format == "html" {
		// Marshal data into pretty JSON for better readability.
		prettyJSON, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			return err
		}

		// A simple HTML template that displays the JSON data inside a <pre> element,
		// with a blue background, white text, and a white-bordered, rounded-corner box.
		const tmpl = `
		<!DOCTYPE html>
		<html>
		<head>
			<meta charset="utf-8">
			<title>Country Info Response</title>
			<style>
				body {
					font-family: sans-serif;
					margin: 20px;
					background-color: #0000cc; /* or any shade of blue you prefer */
					color: #ffffff; /* white text for the page */
				}
				h1 {
					color: #ffffff; /* ensure the heading text is white */
				}
				pre {
					background-color: #ffffff;  /* white background inside the box */
					color: #000000;             /* black text for the JSON content */
					padding: 10px;
					border: 2px solid #ffffff;  /* white border */
					border-radius: 10px;        /* rounded corners */
					white-space: pre-wrap;      /* wrap lines automatically */
					word-wrap: break-word;      /* break long words if needed */
				}
			</style>
		</head>
		<body>
			<h1>Country Info Response</h1>
			<pre>{{.}}</pre>
		</body>
		</html>
		`
		t, err := template.New("jsonPage").Parse(tmpl)
		if err != nil {
			return err
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		return t.Execute(w, string(prettyJSON))
	}

	// If format != html, return plain JSON as usual.
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(data)
}
