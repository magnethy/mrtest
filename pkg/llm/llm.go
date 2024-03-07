// Package llm integrates with LLM services, providing chatbot functionality
package llm

import (
	"rmcode/pkg/config"
	"rmcode/pkg/http_client"
)

type LLMService interface {
	Chat(msg string) (string, error)
}

func NewOpenAILLMService(cfg *config.Config, httpClient http_client.HttpClient) LLMService {
	return &openAILLMService{HttpClient: httpClient, Config: cfg}
}
