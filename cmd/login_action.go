package cmd

import (
	"fmt"

	"atlas/auth"
)

func runLogin() {
	fb := auth.FirebaseAuth{
		URL:    ctx.Base.Identity.Firebase.URL,
		APIKey: ctx.Creds.Firebase.APIKey,
	}

	token, err := auth.ResolveToken(
		ctx.Creds.Firebase.APIKey,
		func() (*auth.Token, error) {
			return fb.Login(
				ctx.Creds.Firebase.Email,
				ctx.Creds.Firebase.Password,
			)
		},
	)

	if err != nil {
		fmt.Println("Login failed:", err)
		return
	}

	fmt.Println("✅ Authenticated")
	fmt.Println("Token expires at:", token.Expiry)
}
