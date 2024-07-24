package client

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// OpenWeatherMapClient is the client for interacting with the OpenWeatheMap API.
type OpenWeatherMapClient struct {
	APIKey     string
	HTTPClient *http.Client
}

// NewClient creates a new instance of OpenWeatherMapClient.
func NewClient(apiKey string) *OpenWeatherMapClient {
	return &OpenWeatherMapClient{
		APIKey:     apiKey,
		HTTPClient: &http.Client{},
	}
}

// GetCurrentWeather retrieves the current weather data for a given city.
func (c *OpenWeatherMapClient) GetCurrentWeather(city string) (map[string]interface{}, error) {
	url := fmt.Sprintf("%sweather?q=%s&appid=%s&units=%s", BaseURL, city, c.APIKey, Units)
	resp, err := c.HTTPClient.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var apiErr APIError
		if err := json.NewDecoder(resp.Body).Decode(&apiErr); err != nil {
			return nil, fmt.Errorf("error decoding API error response: %v", err)
		}
		apiErr.Code = resp.StatusCode
		return nil, &apiErr
	}
	var weatherData map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&weatherData); err != nil {
		return nil, err
	}

	return weatherData, nil
}

// GetWeatherByCoords retrieves the current weather data for a given latitude and longitude.
func (c *OpenWeatherMapClient) GetWeatherByCoords(lat, lon float64) (map[string]interface{}, error) {
	url := fmt.Sprintf("%sweather?lat=%f&lon=%f&appid=%s&units=%s", BaseURL, lat, lon, c.APIKey, Units)
	resp, err := c.HTTPClient.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		var apiErr APIError
		if err := json.NewDecoder(resp.Body).Decode(&apiErr); err != nil {
			return nil, fmt.Errorf("error decoding API error response: %v", err)
		}
		apiErr.Code = resp.StatusCode
		return nil, &apiErr
	}

	var weatherData map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&weatherData); err != nil {
		return nil, err
	}
	return weatherData, nil
}
