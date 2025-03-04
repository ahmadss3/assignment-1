# Assignment-1 Cloud Technology

This project provides a simple Go-based REST web application that offers information about countries and their historical population data. It also includes a status/diagnostics endpoint. Additionally, there is an option to return responses in a basic HTML format instead of plain JSON, for a more visually appealing browser experience.

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
- REST Countries API 
    - Endpoint: http://129.241.150.113:8080/v3.1/
    - Documentation: http://129.241.150.113:8080/
- CountriesNow API 
     - Endpoint:      http://129.241.150.113:3500/api/v0.1/
     - Documentation: https://documenter.getpostman.com/view/1134062/T1LJjU52

The service combines information from these external APIs to provide:
1. General country info, including continents, population, languages, borders, capital, flag, and a list of cities.
   At this endpoint:
   ```https://assignment-1-x0um.onrender.com/countryinfo/v1/info/?format=html```
2. Historical population data for a country, along with an optional year range filter and a calculated mean value.
   At this endpoint:
   ```https://assignment-1-x0um.onrender.com/countryinfo/v1/population/?format=html```
3. A status endpoint that reports the availability of the external APIs and the uptime of this service.
   At this endpoint:
   ```https://assignment-1-x0um.onrender.com/countryinfo/v1/status/?format=html```

---
# How to use:
---
## As a service:
- Please visit the following link to use the service:
- ```https://assignment-1-x0um.onrender.com```

## Endpoints

### 1. `/countryinfo/v1/info/{country_code}?limit={integer}&format=html`
- **Method:** GET  
- **Description:** Returns general country information for the specified two-letter country code (ISO 3166-2).  
- **Query Parameter (optional):** `limit` (integer) to limit the number of returned cities.
- **Query Parameter (optional):** `format`(html) to get a HTML view.

- **Examples:**  
  - `GET /countryinfo/v1/info/NO`  
  - `GET /countryinfo/v1/info/NO?limit=10&format=html`
  **Response**
  - Content type: application/json
  **Body:**
  ``` {
  "name": "Norway",
  "continents": [
    "Europe"
  ],
  "population": 5379475,
  "languages": {
    "nno": "Norwegian Nynorsk",
    "nob": "Norwegian Bokmål",
    "smi": "Sami"
  },
  "borders": [
    "FIN",
    "SWE",
    "RUS"
  ],
  "flag": "https://flagcdn.com/w320/no.png",
  "capital": "Oslo",
  "cities": [
    "Abelvaer",
    "Adalsbruk",
    "Adland",
    "Agdenes",
    "Agotnes"
  ]
}```

### 2. `/countryinfo/v1/population/{country_code}?format=html`
- **Method:** GET  
- **Description:** Returns historical population data for the specified two-letter country code, along with the mean population.  
- **Query Parameter (optional):** `limit={startYear-endYear}` to filter results by year range. 
- **Query Parameter (optional):** `format`(html) to get a HTML view. 
- **Examples:**  
  - `GET /countryinfo/v1/population/NO`  
  - `GET /countryinfo/v1/population/NO?limit=2013-2015&format=html`
  **Response**
  - Content type: application/json
  **Body:**
  ```{
  "mean": 5135154,
  "values": [
    {
      "year": 2013,
      "value": 5079623
    },
    {
      "year": 2014,
      "value": 5137232
    },
    {
      "year": 2015,
      "value": 5188607
    }
  ]
}```

### 3. `/countryinfo/v1/status/`
- **Method:** GET  
- **Description:** Returns diagnostic information about this service, including uptime and the status codes for the external APIs.
- **Query Parameter (optional):** `format`(html) to get a HTML view.  
- **Example:**  
  - `GET /countryinfo/v1/status/?format=html`
  **Response**
  - Content type: application/json
  **Body:**
  ``` {
  "countriesnowapi": 200,
  "restcountriesapi": 200,
  "version": "v1",
  "uptime": 1370
}```

---

## HTML or JSON Output
By default, all endpoints return JSON (`Content-Type: application/json`).  
If you want a simple HTML view (with a blue background), you can add `?format=html` to any endpoint URL. For example:
- `GET /countryinfo/v1/info/NO?format=html`

This feature makes it easier to read the response in a browser. It does not change the underlying data — only how it is displayed.

---

## As a code
### Gitlab
- ` If you have access to the repository, you can clone it from the following link:`
``` https://git.gvk.idi.ntnu.no/course/prog2005/prog2005-2025-workspace/ahmad/assignment-1.git```
### Running
- Please run the code to host the service on your local machine.

**Access the service** at [http://localhost:8080/](http://localhost:8080/) in your browser or through a REST client.

---

## Project Structure

- **main.go**: Entry point, sets up the server and registers endpoints  
- **handlers.go**: Contains HTTP handler functions for each endpoint  
- **services.go**: External API integrations and in-memory data  
- **models.go**: Data models and response structures  
- **utils.go**: Utility functions, including `respondJSONOrHTML`  
- **README.md**: This file

---
## Enjoy using the Country Information Service!
