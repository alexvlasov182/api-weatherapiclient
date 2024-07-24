package main

import (
	"api-weatherapiclient/client"
	"fmt"
	"log"
	"os"
)

func main() {
	apiKey := os.Getenv("OPENWEATHER_API_KEY")
	if apiKey == "" {
		log.Fatalf("API key not set")
	}
	c := client.NewClient(apiKey)

	// Get request to fetch current weather for a city
	city := "London"
	weather, err := c.GetCurrentWeather(city)
	if err != nil {
		log.Fatalf("Error getting current weather: %v", err)
	}
	fmt.Println("Current Weather in", city, ":", weather)

	// Get request to fetch current weather by coordinates
	lat, lon := 47.3769, 8.5417 // Coordinates for Zurich
	weatherByCoords, err := c.GetWeatherByCoords(lat, lon)
	if err != nil {
		log.Fatalf("Error getting weather by coordinates: %v", err)
	}
	fmt.Println("Current weather at coordinates:", weatherByCoords)
}
