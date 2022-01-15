package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/suadev/git-sum/config"
)

var (
	setTokenCmd = &cobra.Command{
		Use:   "set-token",
		Short: "Set GitHub access token.",
		Long:  "Set Github Auth Token to increase your access limit.",
		Run: func(cmd *cobra.Command, args []string) {
			setAccessToken()
		},
	}
)

func setAccessToken() {
	var token = ""
	for token == "" {
		fmt.Print("GitHub Access Token: ")
		fmt.Scan(&token)
		config.SetAccessToken(token)
	}
}

func init() {
	rootCmd.AddCommand(setTokenCmd)
}
