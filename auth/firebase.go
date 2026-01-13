package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type FirebaseAuth struct {
	URL    string
	APIKey string
}

type LoginRequest struct {
	Email             string `json:"email"`
	Password          string `json:"password"`
	ReturnSecureToken bool   `json:"returnSecureToken"`
}

type LoginResponse struct {
	IDToken      string `json:"idToken"`
	RefreshToken string `json:"refreshToken"`
	ExpiresIn    string `json:"expiresIn"`
}

func (f *FirebaseAuth) Login(email, password string) (*Token, error) {
	endpoint := fmt.Sprintf("%s?key=%s", f.URL, f.APIKey)
	fmt.Printf(endpoint)
	payload := LoginRequest{
		Email:             email,
		Password:          password,
		ReturnSecureToken: true,
	}

	body, _ := json.Marshal(payload)
	req, _ := http.NewRequest("POST", endpoint, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("firebase auth failed: %s", resp.Status)
	}

	var out LoginResponse
	expires, _ := time.ParseDuration(out.ExpiresIn + "s")

	return &Token{
		IDToken:      out.IDToken,
		RefreshToken: out.RefreshToken,
		Expiry:       time.Now().Add(expires),
	}, nil
}
