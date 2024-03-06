package llm

import (
	"fmt"
	"rmcode/pkg/config"
	"rmcode/pkg/http_client"
)

type OpenAILLMService struct {
	HttpClient http_client.HttpClient
	Config     *config.Config
}

type OpenAiChatResponse2 struct {
	Choices []struct {
		Index   int `json:"index"`
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

func (o *OpenAILLMService) Chat(msg string) (string, error) {
	var responseObj OpenAiChatResponse2

	headers := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", o.Config.OpenAI.Key),
		"Content-Type":  "application/json",
	}

	requestBody := map[string]interface{}{
		"model": "gpt-3.5-turbo",
		"messages": []map[string]interface{}{
			{
				"role":    "system",
				"content": o.Config.OpenAI.Chat.Context,
			},
			{
				"role":    "user",
				"content": msg,
			},
		}}

	err := o.HttpClient.PostJson(o.Config.OpenAI.Chat.URL, headers, &requestBody, &responseObj)
	if err != nil {
		return "", fmt.Errorf("could not get chat from openai: %s", err)
	}

	return responseObj.Choices[0].Message.Content, nil
}
