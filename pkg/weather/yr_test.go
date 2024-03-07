package weather

import (
	"encoding/json"
	"rmcode/pkg/config"
	"rmcode/pkg/http_client"
	"testing"
)

var cfg *config.Config

func TestMain(m *testing.M) {
	cfg, _ = config.ReadConfig("../../configs/test_config.json")
	m.Run()
}

func TestYr_ShouldCallCorrectEndpoint(t *testing.T) {
	httpMockClient := http_client.MockHttpClient{
		GetFunc: func(url string, headers map[string]string, v any) {
			if url != cfg.Yr.LocationForecast.Url {
				t.Errorf("Expected yr URL to be %s but was %s", cfg.Yr.LocationForecast.Url, url)
			}
		},
	}

	yrService := yrWeatherService{HttpClient: &httpMockClient, Config: cfg}
	yrService.getYrData()
}

func TestYr_ShouldReturnCorrectDataFromAPIResponse(t *testing.T) {
	httpMockClient := http_client.MockHttpClient{
		GetFunc: func(url string, headers map[string]string, v any) {
			mockJson := ` {"properties": {"timeseries": [{"data": {"instant": {"details": {"air_temperature": 15.0}}}}]}} `
			json.Unmarshal([]byte(mockJson), &v)
		},
	}
	yrService := yrWeatherService{HttpClient: &httpMockClient, Config: cfg}

	var expectedAirTemp float32 = 15.0
	w, _ := yrService.GetWeather()

	if w.AirTemperature != expectedAirTemp {
		t.Errorf("Expected the returned object to have air temperature %.2f, but was %.2f", expectedAirTemp, w.AirTemperature)
	}
}
