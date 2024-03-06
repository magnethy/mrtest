package main

import (
	"fmt"
	"rmcode/cmd/rm"
	"rmcode/pkg/config"
	"rmcode/pkg/http_client"
	"rmcode/pkg/llm"
	"rmcode/pkg/weather"
)

func main() {
	cfg, err := config.ReadConfig("configs/config.json")
	if err != nil {
		fmt.Printf("Could not read config: %s\n", err)
		return
	}

	httpClient := http_client.NewJSONHttpClient()
	weatherService := weather.NewYrWeatherService(cfg, httpClient)
	llmService := llm.NewOpenAILLMService(cfg, httpClient)

	rm.Main(weatherService, llmService)
}
