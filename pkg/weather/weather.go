package weather

import (
	"rmcode/pkg/config"
	"rmcode/pkg/http_client"
)

type WeatherService interface {
	GetWeather() (*Weather, error)
}
type Weather struct {
	AirTemperature float32
}

func NewYrWeatherService(cfg *config.Config, httpClient http_client.HttpClient) WeatherService {
	return &yrWeatherService{HttpClient: httpClient, Config: cfg}
}
