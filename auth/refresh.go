package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type refreshResponse struct {
	IDToken      string `json:"id_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    string `json:"expires_in"`
}

func RefreshToken(apiKey, refreshToken string) (*Token, error) {
	fmt.Printf("Refreshing token...\n")
	endpoint := fmt.Sprintf(
		"https://securetoken.googleapis.com/v1/token?key=%s",
		apiKey,
	)

	data := url.Values{}
	data.Set("grant_type", "refresh_token")
	data.Set("refresh_token", refreshToken)

	req, _ := http.NewRequest(
		"POST",
		endpoint,
		strings.NewReader(data.Encode()),
	)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("token refresh failed: %s", resp.Status)
	}

	var out refreshResponse
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return nil, err
	}

	expires, _ := time.ParseDuration(out.ExpiresIn + "s")

	return &Token{
		IDToken:      out.IDToken,
		RefreshToken: out.RefreshToken,
		Expiry:       time.Now().Add(expires),
	}, nil
}
