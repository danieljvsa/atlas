package auth

import (
	"fmt"
	"net/http"
)

// APIClient wraps an HTTP client and adds Authorization header automatically
type APIClient struct {
	BaseURL string
	Token   *Token
	APIKey  string
	Client  *http.Client
}

// NewAPIClient returns a new client
func NewAPIClient(baseURL string, token *Token, apiKey string) *APIClient {
	return &APIClient{
		BaseURL: baseURL,
		Token:   token,
		APIKey:  apiKey,
		Client:  &http.Client{},
	}
}

// DoRequest sends an HTTP request with Authorization header
func (c *APIClient) DoRequest(method, path string, body *http.Request) (*http.Response, error) {
	// Refresh token if needed
	if !c.Token.Valid() {
		nt, err := RefreshToken(c.APIKey, c.Token.RefreshToken)
		if err != nil {
			return nil, fmt.Errorf("token refresh failed: %w", err)
		}
		c.Token = nt
		_ = SaveToken(nt)
	}

	// Create request if not provided
	req := body
	if req == nil {
		var err error
		req, err = http.NewRequest(method, c.BaseURL+path, nil)
		if err != nil {
			return nil, err
		}
	}

	// Add Authorization header
	req.Header.Set("Authorization", "Bearer "+c.Token.IDToken)
	return c.Client.Do(req)
}
