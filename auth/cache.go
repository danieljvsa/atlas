package auth

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

func tokenPath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".config", "atlas", "token.yaml")
}

func LoadToken() (*Token, error) {
	data, err := os.ReadFile(tokenPath())
	if err != nil {
		return nil, err
	}

	var t Token
	if err := yaml.Unmarshal(data, &t); err != nil {
		return nil, err
	}

	return &t, nil
}

func SaveToken(t *Token) error {
	path := tokenPath()
	os.MkdirAll(filepath.Dir(path), 0700)

	data, _ := yaml.Marshal(t)
	return os.WriteFile(path, data, 0600)
}
