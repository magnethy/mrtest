package http_client

// MockHttpClient is used for mocking the HTTP client during testing
type MockHttpClient struct {
	GetFunc  func(url string, headers map[string]string, v any)
	PostFunc func(url string, headers map[string]string, body any, v any)
}

func (h *MockHttpClient) Get(url string, headers map[string]string, v any) error {
	h.GetFunc(url, headers, v)
	return nil
}

func (h *MockHttpClient) Post(url string, headers map[string]string, body any, v any) error {
	h.PostFunc(url, headers, body, v)
	return nil
}
