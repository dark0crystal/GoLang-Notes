package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
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

	fetchWeather("25.28", "55.30")
}
