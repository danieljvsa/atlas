package cmd

import (
	"atlas/config"

	"github.com/AlecAivazis/survey/v2"
)

func promptCredentials() (*config.Credentials, error) {
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

	return &config.Credentials{
		Firebase: config.FirebaseCredentials{
			APIKey:   apiKey,
			Email:    email,
			Password: password,
		},
	}, nil
}
