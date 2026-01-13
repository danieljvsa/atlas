package config

import (
	"errors"
	"os"
	"path/filepath"
)

func ResolveBaseConfig() (string, error) {
	if p := os.Getenv("ATLAS_CONFIG"); p != "" {
		return p, nil
	}

	if _, err := os.Stat("config.yaml"); err == nil {
		return "config.yaml", nil
	}

	home, _ := os.UserHomeDir()
	path := filepath.Join(home, ".config", "atlas", "config.yaml")
	if _, err := os.Stat(path); err == nil {
		return path, nil
	}

	return "", errors.New("base config not found")
}

func ResolveCredentials() (string, error) {
	if p := os.Getenv("ATLAS_CREDS"); p != "" {
		return p, nil
	}

	home, _ := os.UserHomeDir()
	path := filepath.Join(home, ".config", "atlas", "credentials.yaml")
	if _, err := os.Stat(path); err == nil {
		return path, nil
	}

	return "", errors.New("credentials not found")
}
