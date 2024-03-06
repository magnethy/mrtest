package rm

import (
	"fmt"
	"rmcode/pkg/http_client"
	"rmcode/pkg/weather"
)

func Main() {
	httpClient := http_client.NewBasicHttpClient()
	weatherService := weather.NewYrWeatherService(httpClient)

	w, err := weatherService.GetWeather()
	if err != nil {
		fmt.Printf("Got error %s", err)
	}
	fmt.Printf("Weather: %f", w.AirTemperature)
}
