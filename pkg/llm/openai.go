package llm

import (
	"fmt"
	"rmcode/pkg/config"
	"rmcode/pkg/http_client"
)

type openAILLMService struct {
	HttpClient http_client.HttpClient
	Config     *config.Config
}

type openAiChatResponse struct {
	Choices []struct {
		Index   int `json:"index"`
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

func (o *openAILLMService) Chat(msg string) (string, error) {
	var responseObj openAiChatResponse

	headers := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", o.Config.OpenAI.Key),
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

	err := o.HttpClient.Post(o.Config.OpenAI.Chat.URL, headers, &requestBody, &responseObj)
	if err != nil {
		return "", fmt.Errorf("could not get chat from openai: %s", err)
	}

	return responseObj.Choices[0].Message.Content, nil
}
