package weather

import "rmcode/pkg/http_client"

type WeatherService interface {
	GetWeather() (*Weather, error)
}
type Weather struct {
	AirTemperature float32
}

func NewYrWeatherService(httpClient http_client.HttpClient) WeatherService {
	return &YrWeatherService{HttpClient: httpClient}
}
