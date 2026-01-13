package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"atlas/config"

	"github.com/AlecAivazis/survey/v2"
	"gopkg.in/yaml.v3"
)

func maybeSaveCredentials(creds *config.Credentials) {
	var save bool
	survey.AskOne(&survey.Confirm{
		Message: "Save these credentials?",
		Default: false,
	}, &save)

	if !save {
		return
	}

	if creds.Firebase.APIKey == "" || creds.Firebase.Email == "" || creds.Firebase.Password == "" {
		fmt.Println("⚠ Not saving empty credentials.")
		return
	}

	home, _ := os.UserHomeDir()
	dir := filepath.Join(home, ".config", "atlas")
	os.MkdirAll(dir, 0700)

	path := filepath.Join(dir, "credentials.yaml")
	data, _ := yaml.Marshal(creds)

	_ = os.WriteFile(path, data, 0600)
}
