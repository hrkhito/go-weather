package main

import (
	"log"
	"net/http"

	"github.com/hrk_hito/go-weather/handlers"

	"github.com/rs/cors"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/current-weather", handlers.GetCurrentWeather)
	mux.HandleFunc("/forecast", handlers.GetForecast)

	handler := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete},
		AllowedHeaders: []string{"*"},
	}).Handler(mux)

	log.Fatal(http.ListenAndServe(":8080", handler))
}
