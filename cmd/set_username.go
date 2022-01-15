package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/suadev/git-sum/config"
)

var (
	setUserCmd = &cobra.Command{
		Use:   "set-user",
		Short: "Set GitHub Username.",
		Long:  "Change current user to show GitHub summary.",
		Run: func(cmd *cobra.Command, args []string) {
			setUsernameAndPrintSummary()
		},
	}
)

func setUsernameAndPrintSummary() {
	var userName = ""
	for userName == "" {
		fmt.Print("GitHub Username: ")
		fmt.Scan(&userName)
		config.SetUserName(userName)
	}
	userName, authToken := config.GetUserNameAndAuthToken()
	printGithubSummary(userName, authToken)
}

func init() {
	rootCmd.AddCommand(setUserCmd)
}
