package cmd

import (
	"errors"
	"os"

	"atlas/config"

	"github.com/AlecAivazis/survey/v2"
)

func loadCredentialsFromFile() (*config.Credentials, error) {
	defaultPath, _ := config.ResolveCredentials()

	var path string
	survey.AskOne(&survey.Input{
		Message: "Path to credentials file:",
		Default: defaultPath,
	}, &path)

	if _, err := os.Stat(path); err != nil {
		return nil, errors.New("credentials file not found")
	}

	return config.LoadCredentials(path)
}
