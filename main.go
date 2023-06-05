package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type WeatherResponse struct {
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
}

type ForecastResponse struct {
	List []struct {
		Main struct {
			Temp float64 `json:"temp"`
		} `json:"main"`
		Weather []struct {
			Description string `json:"description"`
		} `json:"weather"`
		DtTxt string `json:"dt_txt"`
	} `json:"list"`
}

func main() {
	fmt.Println("Enter the city name:")
	var city string
	fmt.Scanln(&city)

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey := os.Getenv("OPENWEATHER_API_KEY")

	response, err := http.Get(fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, apiKey))

	if err != nil {
		fmt.Println("Error getting response from API:", err)
		return
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	var weatherResponse WeatherResponse
	err = json.Unmarshal(responseData, &weatherResponse)
	if err != nil {
		fmt.Println("Error unmarshalling response:", err)
		return
	}

	if weatherResponse.Main.Temp != 0 && len(weatherResponse.Weather) > 0 {
		celsius := weatherResponse.Main.Temp - 273.15
		fmt.Println("Current Temperature (Celsius): ", celsius)
		fmt.Println("Current Weather Condition : ", weatherResponse.Weather[0].Description)
	} else {
		fmt.Println("No weather data available for this location")
	}

	forecastResponse, err := http.Get(fmt.Sprintf("http://api.openweathermap.org/data/2.5/forecast?q=%s&appid=%s", city, apiKey))
	if err != nil {
		fmt.Println("Error getting response from API:", err)
		return
	}

	forecastData, err := io.ReadAll(forecastResponse.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	var forecast ForecastResponse
	err = json.Unmarshal(forecastData, &forecast)
	if err != nil {
		fmt.Println("Error unmarshalling response:", err)
		return
	}

	for _, item := range forecast.List {
		celsius := item.Main.Temp - 273.15
		fmt.Println("Time: ", item.DtTxt, " Temperature (Celsius): ", celsius, " Weather: ", item.Weather[0].Description)
	}
}
