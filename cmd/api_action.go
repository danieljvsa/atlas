package cmd

import (
	"atlas/api"
	"atlas/auth"
	"fmt"
	"io/ioutil"

	"github.com/AlecAivazis/survey/v2"
)

func runAPICall() {
	// Ask user which endpoint
	var path string
	survey.AskOne(&survey.Input{
		Message: "Enter API path (e.g., /users/list):",
		Default: "/",
	}, &path)

	client := auth.NewAPIClient(
		ctx.Base.Environments[ctx.Env].BaseURL,
		nil, // token will be loaded via resolver
		ctx.Creds.Firebase.APIKey,
	)

	// Resolve token automatically
	token, err := auth.ResolveToken(
		ctx.Creds.Firebase.APIKey,
		func() (*auth.Token, error) {
			fb := auth.FirebaseAuth{
				URL:    ctx.Base.Identity.Firebase.URL,
				APIKey: ctx.Creds.Firebase.APIKey,
			}
			return fb.Login(ctx.Creds.Firebase.Email, ctx.Creds.Firebase.Password)
		},
	)
	if err != nil {
		fmt.Println("Failed to get token:", err)
		return
	}
	client.Token = token

	resp, err := client.DoRequest("GET", path, nil)
	if err != nil {
		fmt.Println("API request failed:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Status:", resp.Status)
	fmt.Println("Body:", string(body))
}

func handleAPIInteraction() {
	path := "./api_actions.yaml"
	af, err := api.LoadActions(path)
	if err != nil {
		fmt.Println("Failed to load API actions:", err)
		return
	}

	client := auth.NewAPIClient(
		ctx.Base.Environments[ctx.Env].BaseURL,
		nil,
		ctx.Creds.Firebase.APIKey,
	)

	token, err := auth.ResolveToken(
		ctx.Creds.Firebase.APIKey,
		func() (*auth.Token, error) {
			fb := auth.FirebaseAuth{
				URL:    ctx.Base.Identity.Firebase.URL,
				APIKey: ctx.Creds.Firebase.APIKey,
			}
			return fb.Login(ctx.Creds.Firebase.Email, ctx.Creds.Firebase.Password)
		},
	)
	if err != nil {
		fmt.Println("Failed to get token:", err)
		return
	}
	client.Token = token

	api.ExecuteActions(client, af)
}
