package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func LoadBase(path string) (*BaseConfig, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg BaseConfig
	return &cfg, yaml.Unmarshal(data, &cfg)
}

func LoadCredentials(path string) (*Credentials, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	fmt.Printf(string(data))

	var creds Credentials
	return &creds, yaml.Unmarshal(data, &creds)
}
