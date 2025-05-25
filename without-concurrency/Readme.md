🌦 Weather Fetcher in Go
This Go program fetches the current temperature from the OpenWeatherMap API for a given latitude and longitude using an API key stored in an environment variable.


🧪 Run the Program

go run main.go


📌 fetchWeather(lat, lon string)
This function:

Reads the API key from the environment using os.Getenv("WeatherOpenApiKey").

Constructs the API request URL using the given latitude and longitude.

Sends an HTTP GET request to the OpenWeatherMap API.

Parses the response JSON to extract the current temperature.

Prints the temperature to the console.

📦 JSON Structure Extracted

{
  "main": {
    "temp": 34.5
  }
}


📌 main()
This function:

Loads the .env file from the parent directory (../.env).

Calls fetchWeather() with latitude 25.28 and longitude 55.30.