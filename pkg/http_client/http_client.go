// Package http_client is a reusable tool for doing HTTP requests without any extra clutter
package http_client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type HttpClient interface {
	GetJson(url string, headers map[string]string, v any) error
}

type BasicHttpClient struct{}

func NewBasicHttpClient() HttpClient {
	return &BasicHttpClient{}
}

// GetJson fetches a URL and fills the `v` parameter with the JSON response
func (h *BasicHttpClient) GetJson(url string, headers map[string]string, v any) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("could not create request, %s", err)
	}

	for headerKey, headerValue := range headers {
		req.Header.Add(headerKey, headerValue)
	}

	client := http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("could not get response, %s", err)
	}

	defer response.Body.Close()

	jsonBody, err := io.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("error read response body, %s", err)
	}

	err = json.Unmarshal(jsonBody, &v)

	if err != nil {
		return fmt.Errorf("could not convert response body to json, %s", err)
	}

	return nil
}
