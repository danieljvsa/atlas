package cmd

import (
	"os"
	"path/filepath"

	"atlas/config"

	"github.com/AlecAivazis/survey/v2"
	"gopkg.in/yaml.v3"
)

func setupCredentials() error {
	var apiKey, email, password string

	survey.AskOne(&survey.Input{
		Message: "Firebase API key:",
	}, &apiKey)

	survey.AskOne(&survey.Input{
		Message: "Email:",
	}, &email)

	survey.AskOne(&survey.Password{
		Message: "Password:",
	}, &password)

	creds := config.Credentials{
		Firebase: config.FirebaseCredentials{
			APIKey:   apiKey,
			Email:    email,
			Password: password,
		},
	}

	home, _ := os.UserHomeDir()
	dir := filepath.Join(home, ".config", "atlas")
	os.MkdirAll(dir, 0700)

	data, _ := yaml.Marshal(&creds)
	return os.WriteFile(filepath.Join(dir, "credentials.yaml"), data, 0600)
}
