// Script to geocode providers without coordinates
package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

// GeocodingResponse represents the response from Google's Geocoding API
type GeocodingResponse struct {
	Results []struct {
		Geometry struct {
			Location struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"location"`
		} `json:"geometry"`
	} `json:"results"`
	Status string `json:"status"`
}

func main() {
	// Connect to database
	connStr := "postgresql://yourusername:yourpassword@localhost/yourdatabase?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	defer db.Close()

	// Get API key from environment variable
	apiKey := os.Getenv("GOOGLE_MAPS_API_KEY")
	if apiKey == "" {
		log.Fatal("GOOGLE_MAPS_API_KEY environment variable is required")
	}

	// Get providers without coordinates
	rows, err := db.Query(`
		SELECT id, name, address 
		FROM providers 
		WHERE latitude = 0 OR longitude = 0 OR latitude IS NULL OR longitude IS NULL
	`)
	if err != nil {
		log.Fatal("Error querying providers:", err)
	}
	defer rows.Close()

	// Process each provider
	for rows.Next() {
		var id int
		var name, address string
		var addressToGeocode string

		if err := rows.Scan(&id, &name, &address); err != nil {
			log.Printf("Error scanning row: %v", err)
			continue
		}

		// Determine what to geocode
		if address != "" {
			addressToGeocode = address
		} else {
			// Use name + Los Angeles if no address is available
			addressToGeocode = name + ", Los Angeles, CA"
		}

		// Geocode the address
		log.Printf("Geocoding provider #%d: %s", id, name)
		lat, lng, err := geocodeAddress(addressToGeocode, apiKey)
		if err != nil {
			log.Printf("Error geocoding address for provider #%d: %v", id, err)
			continue
		}

		// Update the database with the coordinates
		_, err = db.Exec(`
			UPDATE providers 
			SET latitude = $1, longitude = $2 
			WHERE id = $3
		`, lat, lng, id)
		if err != nil {
			log.Printf("Error updating coordinates for provider #%d: %v", id, err)
			continue
		}

		log.Printf("Successfully updated coordinates for provider #%d: %f, %f", id, lat, lng)

		// Sleep to avoid hitting API rate limits
		time.Sleep(200 * time.Millisecond)
	}

	log.Println("Geocoding process completed")
}

// geocodeAddress converts an address to latitude and longitude using Google's Geocoding API
func geocodeAddress(address string, apiKey string) (float64, float64, error) {
	// URL encode the address
	escapedAddress := url.QueryEscape(address)

	// Build the API URL
	apiURL := fmt.Sprintf(
		"https://maps.googleapis.com/maps/api/geocode/json?address=%s&key=%s&components=country:US",
		escapedAddress,
		apiKey,
	)

	// Make the request
	resp, err := http.Get(apiURL)
	if err != nil {
		return 0, 0, fmt.Errorf("error making geocoding request: %v", err)
	}
	defer resp.Body.Close()

	// Parse the response
	var result GeocodingResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, 0, fmt.Errorf("error decoding geocoding response: %v", err)
	}

	// Check if the request was successful
	if result.Status != "OK" {
		return 0, 0, fmt.Errorf("geocoding API returned status: %s", result.Status)
	}

	// Check if we got any results
	if len(result.Results) == 0 {
		return 0, 0, fmt.Errorf("no geocoding results found for address: %s", address)
	}

	// Return the coordinates
	lat := result.Results[0].Geometry.Location.Lat
	lng := result.Results[0].Geometry.Location.Lng
	return lat, lng, nil
}