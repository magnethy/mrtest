// Package weather connects to yr API for weather data
package weather

import (
	"fmt"
	"rmcode/pkg/http_client"
)

type YrWeatherService struct {
	HttpClient http_client.HttpClient
}

func (s *YrWeatherService) GetWeather() (*Weather, error) {
	yrData, err := s.getYrData()
	if err != nil {
		return &Weather{}, fmt.Errorf("got error from yr: %s", err)
	}

	weather := new(Weather)
	weather.AirTemperature = yrData.Properties.Timeseries[0].Data.Instant.Details.AirTemperature
	return weather, nil
}

// getYrData fetches weather data for a fixed coordinate
func (s *YrWeatherService) getYrData() (YrLocationForecast, error) {
	apiUrl := "https://api.met.no/weatherapi/locationforecast/2.0/compact?lat=60.10&lon=9.58"

	var resObj YrLocationForecast

	headers := map[string]string{
		"User-Agent": "magne@thyrhaug.net",
	}

	err := s.HttpClient.GetJson(apiUrl, headers, &resObj)
	if err != nil {
		return YrLocationForecast{}, err
	}

	return resObj, nil
}
