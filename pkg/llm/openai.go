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

type openAIChatRequestBodyMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
type openAIChatRequestBody struct {
	Model    string                         `json:"model"`
	Messages []openAIChatRequestBodyMessage `json:"messages"`
}

func (o *openAILLMService) Chat(msg string) (string, error) {
	var responseObj openAiChatResponse

	headers := map[string]string{
		"Authorization": fmt.Sprintf("Bearer %s", o.Config.OpenAI.Key),
	}

	reqBody := openAIChatRequestBody{
		Model: "gpt-3.5-turbo",
		Messages: []openAIChatRequestBodyMessage{
			{Role: "system", Content: o.Config.OpenAI.Chat.Context},
			{Role: "user", Content: msg},
		},
	}

	err := o.HttpClient.Post(o.Config.OpenAI.Chat.URL, headers, &reqBody, &responseObj)
	if err != nil {
		return "", fmt.Errorf("could not get chat from openai: %s", err)
	}

	return responseObj.Choices[0].Message.Content, nil
}
