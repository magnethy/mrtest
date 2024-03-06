// Package config reads json configuration file(s), for the app to use
package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Yr struct {
		LocationForecast struct {
			Url string `json:"url"`
		} `json:"locationforecast"`
		User string `json:"user"`
	} `json:"yr"`
	OpenAI struct {
		Key  string `json:"key"`
		Chat struct {
			URL     string `json:"url"`
			Context string `json:"context"`
		} `json:"chat"`
	} `json:"openai"`
}

// ReadConfig parses the JSON file at the path and returns a struct containing the config values
func ReadConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return &Config{}, fmt.Errorf("could not open config file: %s", err)
	}

	var config Config
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return &Config{}, fmt.Errorf("could not parse json of config file: %s", err)
	}

	return &config, nil
}
