package rm

import (
	"fmt"
	"rmcode/pkg/llm"
	"rmcode/pkg/weather"
)

func Main(weatherService weather.WeatherService, llmService llm.LLMService) {
	w, err := weatherService.GetWeather()
	if err != nil {
		fmt.Printf("Got weather error %s", err)
		return
	}

	llmResponse, err := llmService.Chat(fmt.Sprintf("Talk about the weather being %.1f celsius", w.AirTemperature))
	if err != nil {
		fmt.Printf("Got LLM error: %s", err)
		return
	}

	fmt.Println(llmResponse)
}
