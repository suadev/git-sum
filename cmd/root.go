package cmd

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/suadev/git-sum/config"
	"github.com/suadev/git-sum/github"
)

var (
	rootCmd = &cobra.Command{
		Use:  "git-sum",
		Long: "git-sum shows your current issue and pr counts for each repository of given user.",
	}
)

func Execute() {
	rootCmd.Run = func(cmd *cobra.Command, args []string) {
		userName, authToken := config.GetUserNameAndAuthToken()
		for userName == "" {
			fmt.Print("GitHub Username: ")
			fmt.Scan(&userName)
			config.SetUserName(userName)
		}
		printGithubSummary(userName, authToken)
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func printGithubSummary(userName string, authToken string) {
	githubClient := github.NewClient()

	printHeader()
	for repoList := range githubClient.GetRepoListAsBatch(userName, authToken) {
		var wg sync.WaitGroup
		wg.Add(len(repoList))

		for _, repo := range repoList {
			r := repo
			go func() {
				defer wg.Done()
				repoDetail := githubClient.GetRepo(r.Url, authToken)
				pullRequestCount := githubClient.GetPullRequestCount(r.Url+"/pulls", authToken)
				if repoDetail == nil || (repoDetail.IssueCount < 1 && pullRequestCount < 1) {
					return
				}
				printLine(repoDetail.Name, repoDetail.IssueCount, pullRequestCount)
			}()
		}
		wg.Wait()
		color.Unset()
	}
}

func printHeader() {
	color.Blue("REPO" + getSpaces(60) + "ISSUE COUNT" + getSpaces(10) + "PR COUNT")
}

func printLine(repoName string, issueCount int, pullRequestCount int) {
	firstSpaceCount := 64 - len(repoName)
	issueCountStr := strconv.Itoa(issueCount)
	pullRequestCountStr := strconv.Itoa(pullRequestCount)
	color.Green(repoName + getSpaces(firstSpaceCount) + issueCountStr + getSpaces(21-len(issueCountStr)) + pullRequestCountStr)
}

func getSpaces(count int) string {
	return strings.Repeat(" ", count)
}
