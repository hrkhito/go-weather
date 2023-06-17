package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/hrk_hito/go-weather/models"

	"github.com/joho/godotenv"
)

func GetCurrentWeather(w http.ResponseWriter, r *http.Request) {
	city := r.URL.Query().Get("city")

	err := godotenv.Load()
	if err != nil {
		http.Error(w, "Error loading .env file", http.StatusInternalServerError)
		return
	}

	apiKey := os.Getenv("OPENWEATHER_API_KEY")

	response, err := http.Get(fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, apiKey))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var weatherResponse models.WeatherResponse
	err = json.Unmarshal(responseData, &weatherResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	celsius := weatherResponse.Main.Temp - 273.15
	json.NewEncoder(w).Encode(map[string]interface{}{
		"temperature": celsius,
		"condition":   weatherResponse.Weather[0].Description,
	})
}

func GetForecast(w http.ResponseWriter, r *http.Request) {
	city := r.URL.Query().Get("city")

	err := godotenv.Load()
	if err != nil {
		http.Error(w, "Error loading .env file", http.StatusInternalServerError)
		return
	}

	apiKey := os.Getenv("OPENWEATHER_API_KEY")

	forecastResponse, err := http.Get(fmt.Sprintf("http://api.openweathermap.org/data/2.5/forecast?q=%s&appid=%s", city, apiKey))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer forecastResponse.Body.Close()

	forecastData, err := io.ReadAll(forecastResponse.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var forecast models.ForecastResponse
	err = json.Unmarshal(forecastData, &forecast)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var forecastList []map[string]interface{}

	for _, item := range forecast.List {
		celsius := item.Main.Temp - 273.15
		forecastList = append(forecastList, map[string]interface{}{
			"time":        item.DtTxt,
			"temperature": celsius,
			"condition":   item.Weather[0].Description,
		})
	}
	json.NewEncoder(w).Encode(forecastList)
}
