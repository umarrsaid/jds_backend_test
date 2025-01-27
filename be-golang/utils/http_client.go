package utils

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

// HTTPClient handles common HTTP operations
type HTTPClient struct {
	domain string
	getAuthToken func() (string, error)
}

// NewHTTPClient creates a new HTTP client helper
func NewHTTPClient(domain string, authTokenFn func() (string, error)) *HTTPClient {
	return &HTTPClient{
		domain: domain,
		getAuthToken: authTokenFn,
	}
}

// MakeRequest is a helper function to handle common HTTP request logic
func (h *HTTPClient) MakeRequest(method, endpoint string, body []byte, needsAuth bool) ([]byte, error) {
	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseRequest(req)
	defer fasthttp.ReleaseResponse(resp)

	req.SetRequestURI(fmt.Sprintf("https://%s%s", h.domain, endpoint))
	req.Header.SetMethod(method)
	
	if needsAuth {
		token, err := h.getAuthToken()
		if err != nil {
			return nil, fmt.Errorf("failed to get auth token: %v", err)
		}
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	}

	if body != nil {
		req.Header.SetContentType("application/json")
		req.SetBody(body)
	}

	if err := fasthttp.Do(req, resp); err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}

	if resp.StatusCode() >= 400 {
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode(), resp.Body())
	}

	return resp.Body(), nil
}