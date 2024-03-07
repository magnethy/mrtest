// Package http_client is a reusable tool for doing HTTP requests without any extra clutter
package http_client

type HttpClient interface {
	Get(url string, headers map[string]string, v any) error
	Post(url string, headers map[string]string, body any, v any) error
}

func NewJSONHttpClient() HttpClient {
	return &jsonHttpClient{}
}
