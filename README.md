# Assignment_1 Cloud Technology

This project provides a simple Go-based REST web application that offers information about countries and their historical population data. It also includes a status/diagnostics endpoint. Additionally, there is an option to return responses in a basic HTML format (with a blue background) instead of plain JSON, for a more visually appealing browser experience.

## Table of Contents
1. [Overview](#overview)  
2. [Endpoints](#endpoints)  
3. [HTML or JSON Output](#html-or-json-output)  
4. [How to Run](#how-to-run)  
5. [Example Usage](#example-usage)  
6. [Project Structure](#project-structure)  
7. [Notes and Limitations](#notes-and-limitations)

---

## Overview
The application integrates data from:
- REST Countries API (self-hosted at http://129.241.150.113:8080/v3.1/)
- CountriesNow API (self-hosted at http://129.241.150.113:3500/api/v0.1/)

The service combines information from these external APIs to provide:
1. General country info, including continents, population, languages, borders, capital, flag, and a list of cities.  
2. Historical population data for a country, along with an optional year range filter and a calculated mean value.  
3. A status endpoint that reports the availability of the external APIs and the uptime of this service.

---

## Endpoints

### 1. `/countryinfo/v1/info/{country_code}`
- **Method:** GET  
- **Description:** Returns general country information for the specified two-letter country code (ISO 3166-2).  
- **Query Parameter (optional):** `limit` (integer) to limit the number of returned cities.  
- **Examples:**  
  - `GET /countryinfo/v1/info/NO`  
  - `GET /countryinfo/v1/info/NO?limit=10`

### 2. `/countryinfo/v1/population/{country_code}`
- **Method:** GET  
- **Description:** Returns historical population data for the specified two-letter country code, along with the mean population.  
- **Query Parameter (optional):** `limit={startYear-endYear}` to filter results by year range.  
- **Examples:**  
  - `GET /countryinfo/v1/population/NO`  
  - `GET /countryinfo/v1/population/NO?limit=2010-2015`

### 3. `/countryinfo/v1/status/`
- **Method:** GET  
- **Description:** Returns diagnostic information about this service, including uptime and the status codes for the external APIs.  
- **Example:**  
  - `GET /countryinfo/v1/status/`

---

## HTML or JSON Output
By default, all endpoints return JSON (`Content-Type: application/json`).  
If you want a simple HTML view (with a blue background and a white box for JSON data), you can add `?format=html` to any endpoint URL. For example:
- `GET /countryinfo/v1/info/NO?format=html`

This feature makes it easier to read the response in a browser. It does not change the underlying data—only how it is displayed.

---

## How to Run

1. **Clone the repository** (or copy these files) into your local Go workspace.  
2. **Navigate** to the project directory.  
3. **Build and run**:  
   - `go build && ./country-info-service`  
   - or simply `go run .`  
4. **Access the service** at [http://localhost:8080/](http://localhost:8080/) in your browser or through a REST client.

---

## Example Usage

1. **General Info about Norway (JSON)**  
   - URL: `http://localhost:8080/countryinfo/v1/info/NO`  
   - Response: A JSON object containing details such as `name`, `continents`, `population`, `languages`, `borders`, `flag`, `capital`, and `cities`.

2. **General Info about Norway (HTML)**  
   - URL: `http://localhost:8080/countryinfo/v1/info/NO?format=html`  
   - Response: An HTML page with a blue background and a white box displaying the JSON data.

3. **Population Data (JSON)**  
   - URL: `http://localhost:8080/countryinfo/v1/population/NO`  
   - Response: A JSON object that includes the historical population records and a calculated `mean` value.

4. **Population Data with Year Range Filter (HTML)**  
   - URL: `http://localhost:8080/countryinfo/v1/population/NO?limit=2010-2015&format=html`  
   - Response: An HTML page showing the filtered records (from 2010 to 2015) and the average population for that period.

5. **Status (JSON)**  
   - URL: `http://localhost:8080/countryinfo/v1/status/`  
   - Response: A JSON object with diagnostic details including the status codes for the external APIs, service version, and uptime.

6. **Status (HTML)**  
   - URL: `http://localhost:8080/countryinfo/v1/status/?format=html`  
   - Response: The same information as the JSON response, but displayed in a styled HTML page.

---

## Project Structure

- **main.go**: Entry point, sets up the server and registers endpoints  
- **handlers.go**: Contains HTTP handler functions for each endpoint  
- **services.go**: External API integrations and in-memory data  
- **models.go**: Data models and response structures  
- **utils.go**: Utility functions, including `respondJSONOrHTML`  
- **README.md**: This file

---

## Notes and Limitations

1. **In-memory Population Data**  
   The project uses a simple map (`populationDataMap`) to simulate a database of historical population data. This is primarily for demonstration and testing. You can extend it to support additional countries or replace it with a real database in production.

2. **External API Availability**  
   The REST Countries API and CountriesNow API are self-hosted for this assignment. If they become unreachable or return errors, this service will respond with appropriate HTTP error responses (e.g., 500).

3. **Deployment**  
   In a real-world scenario, you might deploy this service on a platform like Render, AWS, or another cloud provider. You could also add environment variables for configuration, logging, and other production-grade features.

4. **Further Enhancements**  
   - Implement caching to reduce calls to external APIs.  
   - Add more robust error handling and logging.  
   - Create a more advanced front-end or documentation page for easier usage.

This service meets the assignment requirements by:
- Providing general country information via `/info/`.  
- Offering historical population data (including optional year range filtering) via `/population/`.  
- Delivering diagnostic information via `/status/`.  
- Supporting both JSON and optional HTML output for improved readability.

Enjoy using the Country Information Service!
