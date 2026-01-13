package cmd

import (
	"atlas/config"

	"github.com/AlecAivazis/survey/v2"
)

type AppContext struct {
	Base  *config.BaseConfig
	Creds *config.Credentials
	Env   string
}

var ctx *AppContext

func initApp() error {
	basePath, err := config.ResolveBaseConfig()
	if err != nil {
		return err
	}

	base, err := config.LoadBase(basePath)
	if err != nil {
		return err
	}

	creds, err := loadCredentialsFlow()
	if err != nil {
		return err
	}

	env := base.CurrentEnv
	if env == "" {
		env = selectEnvironment(base)
	}

	ctx = &AppContext{
		Base:  base,
		Creds: creds,
		Env:   env,
	}

	return nil
}

func selectEnvironment(cfg *config.BaseConfig) string {
	options := make([]string, 0, len(cfg.Environments))
	for k := range cfg.Environments {
		options = append(options, k)
	}

	var env string
	survey.AskOne(&survey.Select{
		Message: "Select environment:",
		Options: options,
	}, &env)

	return env
}
