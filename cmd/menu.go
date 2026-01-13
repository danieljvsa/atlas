package cmd

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
)

func RunMenu() {
	for {
		var choice string

		survey.AskOne(&survey.Select{
			Message: "What do you want to do?",
			Options: []string{
				"Login (Firebase)",
				"Who am I",
				"Select environment",
				"Call API",
				"Exit",
			},
		}, &choice)

		switch choice {
		case "Login (Firebase)":
			runLogin()

		case "Who am I":
			runWhoAmI()

		case "Select environment":
			ctx.Env = selectEnvironment(ctx.Base)

		case "Call API":
			handleAPIInteraction()

		case "Exit":
			fmt.Println("Goodbye 👋")
			return
		}
	}
}
