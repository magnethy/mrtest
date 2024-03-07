package http_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type jsonHttpClient struct{}

// Get fetches a URL and fills the `v` parameter with the JSON response
func (h *jsonHttpClient) Get(url string, headers map[string]string, v any) error {
	err := doRequest("GET", url, headers, nil, v)
	if err != nil {
		return fmt.Errorf("could not do POST request: %s", err)
	}

	return nil
}

func (h *jsonHttpClient) Post(url string, headers map[string]string, body any, v any) error {
	err := doRequest("POST", url, headers, body, v)
	if err != nil {
		return fmt.Errorf("could not do POST request: %s", err)
	}

	return nil
}

func doRequest(verb string, url string, headers map[string]string, body any, v any) error {
	response, err := submitRequest(verb, url, headers, body, v)
	if err != nil {
		return fmt.Errorf("could not get %s response, %s", verb, err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("could not close response body: %s", err)
		}
	}(response.Body)

	err = getResponseBody(response, &v)
	if err != nil {
		return fmt.Errorf("could not get response body from %s request: %s", verb, err)
	}

	return nil
}

func submitRequest(verb string, url string, headers map[string]string, body any, v any) (*http.Response, error) {
	jsonRequestBody, _ := json.Marshal(body)
	requestbodyBuffer := bytes.NewBuffer(jsonRequestBody)

	req, err := http.NewRequest(verb, url, requestbodyBuffer)
	if err != nil {
		return &http.Response{}, fmt.Errorf("could not create request, %s", err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	for headerKey, headerValue := range headers {
		req.Header.Add(headerKey, headerValue)
	}

	client := http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return &http.Response{}, fmt.Errorf("could not send request: %s", err)
	}

	return response, nil
}

func getResponseBody(response *http.Response, v any) error {
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
