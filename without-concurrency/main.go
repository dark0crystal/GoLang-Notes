package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func fetchWeather(lat, lon string) {
	apiKey := os.Getenv("WeatherOpenApiKey")
	if apiKey == "" {
		fmt.Println("Missing API key in environment variable")
		return
	}

	var data struct {
		Main struct {
			Temp float64 `json:"temp"`
		} `json:"main"`
	}

	url := fmt.Sprintf(
		"https://api.openweathermap.org/data/2.5/weather?lat=%s&lon=%s&appid=%s&units=metric",
		lat, lon, apiKey,
	)

	res, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching weather:", err)
		return
	}
	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		fmt.Println("Error decoding response:", err)
		return
	}

	fmt.Printf("Temperature at (%s, %s): %.2f°C\n", lat, lon, data.Main.Temp)
}

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	// List of locations
	locations := [][2]string{
		{"25.28", "55.30"}, // Dubai
		{"23.58", "58.40"}, // Muscat
		{"21.42", "39.82"}, // Jeddah
	}

	start := time.Now() // Start time

	for _, loc := range locations {
		fetchWeather(loc[0], loc[1])
	}

	elapsed := time.Since(start) // End time

	fmt.Printf("Total time taken: %s\n", elapsed)
}
