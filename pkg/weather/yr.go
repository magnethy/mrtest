// Package weather connects to yr API for weather data
package weather

import (
	"fmt"
	"rmcode/pkg/config"
	"rmcode/pkg/http_client"
)

type yrLocationForecast struct {
	Properties struct {
		Timeseries []struct {
			Time string `json:"time"`
			Data struct {
				Instant struct {
					Details struct {
						AirTemperature float32 `json:"air_temperature"`
					} `json:"details"`
				} `json:"instant"`
			} `json:"data"`
		} `json:"timeseries"`
	} `json:"properties"`
	Type string `json:"type"`
}

type yrWeatherService struct {
	HttpClient http_client.HttpClient
	Config     *config.Config
}

func (s *yrWeatherService) GetWeather() (*Weather, error) {
	yrData, err := s.getYrData()
	if err != nil {
		return &Weather{}, fmt.Errorf("got error from yr: %s", err)
	}

	weather := new(Weather)
	weather.AirTemperature = yrData.Properties.Timeseries[0].Data.Instant.Details.AirTemperature
	return weather, nil
}

// getYrData fetches weather data for a fixed coordinate
func (s *yrWeatherService) getYrData() (yrLocationForecast, error) {
	var resObj yrLocationForecast

	headers := map[string]string{
		"User-Agent": s.Config.Yr.User,
	}

	err := s.HttpClient.Get(s.Config.Yr.LocationForecast.Url, headers, &resObj)
	if err != nil {
		return yrLocationForecast{}, err
	}

	return resObj, nil
}
