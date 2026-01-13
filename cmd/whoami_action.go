package cmd

import "fmt"

func runWhoAmI() {
	if ctx.Creds == nil {
		fmt.Println("Not authenticated")
		return
	}

	fmt.Println("Current environment:", ctx.Env)
	fmt.Println("Firebase email:", ctx.Creds.Firebase.Email)
}
