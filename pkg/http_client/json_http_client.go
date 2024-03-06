package http_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type JSONHttpClient struct{}

// Get fetches a URL and fills the `v` parameter with the JSON response
func (h *JSONHttpClient) Get(url string, headers map[string]string, v any) error {
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

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("could not close response body: %s", err)
		}
	}(response.Body)

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

func (h *JSONHttpClient) Post(url string, headers map[string]string, body any, v any) error {
	jsonRequestBody, _ := json.Marshal(body)
	requestbodyBuffer := bytes.NewBuffer(jsonRequestBody)
	req, err := http.NewRequest("POST", url, requestbodyBuffer)
	if err != nil {
		return fmt.Errorf("could not create POST request, %s", err)
	}

	for headerKey, headerValue := range headers {
		req.Header.Add(headerKey, headerValue)
	}

	client := http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("could not get response, %s", err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("could not close response body: %s", err)
		}
	}(response.Body)

	jsonResponseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("error read response body, %s", err)
	}

	err = json.Unmarshal(jsonResponseBody, &v)

	if err != nil {
		return fmt.Errorf("could not convert response body to json, %s", err)
	}

	return nil
}
