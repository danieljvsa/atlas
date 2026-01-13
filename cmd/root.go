package cmd

import (
	"fmt"
	"os"
)

func Execute() {
	if err := initApp(); err != nil {
		fmt.Println("Startup error:", err)
		os.Exit(1)
	}

	RunMenu()
}
