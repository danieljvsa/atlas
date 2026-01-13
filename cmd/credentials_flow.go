package cmd

import (
	"atlas/config"
	"fmt"

	"github.com/AlecAivazis/survey/v2"
)

func loadCredentialsFlow() (*config.Credentials, error) {
	// 1️⃣ Check if previous credentials exist
	if path, err := config.ResolveCredentials(); err == nil {
		if useLast(path) {
			return config.LoadCredentials(path)
		}
	}

	// 2️⃣ Ask how to load credentials
	var choice string
	survey.AskOne(&survey.Select{
		Message: "Do you want to load credentials from a YAML file?",
		Options: []string{"Yes", "No"},
		Default: "Yes",
	}, &choice)

	if choice == "Yes" {
		return loadCredentialsFromFile()
	}

	// 3️⃣ Manual entry
	creds, err := promptCredentials()
	if err != nil {
		return nil, err
	}

	// Check if any field is empty
	if creds.Firebase.APIKey == "" || creds.Firebase.Email == "" || creds.Firebase.Password == "" {
		fmt.Println("⚠ Credentials file is missing values, switching to manual input.")
		return promptCredentials()
	}

	// 4️⃣ Ask to save
	maybeSaveCredentials(creds)

	return creds, nil
}
