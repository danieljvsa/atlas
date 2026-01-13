package cmd

import "github.com/AlecAivazis/survey/v2"

func useLast(path string) bool {
	var use bool
	survey.AskOne(&survey.Confirm{
		Message: "Use last credentials? (" + path + ")",
		Default: true,
	}, &use)

	return use
}
