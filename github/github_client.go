package github

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type GithubClient struct {
	BaseUrl    string
	HttpClient *http.Client
}

func NewClient() *GithubClient {
	return &GithubClient{
		BaseUrl:    "https://api.github.com",
		HttpClient: &http.Client{},
	}
}

func (c *GithubClient) GetRepoListAsBatch(userName string, authToken string) <-chan []repoListItem {
	repoListChan := make(chan []repoListItem)
	page := 1
	perPage := 100

	go func() {
		for true {
			allRepoEndpoint := fmt.Sprintf("%s/users/%s/repos?page=%d&per_page=%d", c.BaseUrl, userName, page, perPage)
			req, _ := http.NewRequest("GET", allRepoEndpoint, nil)

			addAuthToken(authToken, req)

			var pageResult []repoListItem
			res, err := c.HttpClient.Do(req)
			if err != nil {
				fmt.Println(err)
				repoListChan <- pageResult
				continue
			}

			defer res.Body.Close()

			body, _ := ioutil.ReadAll(res.Body)
			json.Unmarshal(body, &pageResult)

			if len(pageResult) == 0 {
				break
			}
			page++

			repoListChan <- pageResult
		}
		close(repoListChan)
	}()
	return repoListChan
}

func (c *GithubClient) GetRepo(url string, authToken string) *repoItem {
	req, _ := http.NewRequest("GET", url, nil)
	addAuthToken(authToken, req)

	res, err := c.HttpClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var repo repoItem
	json.Unmarshal(body, &repo)
	return &repo
}

func (c *GithubClient) GetPullRequestCount(url string, authToken string) int {
	req, _ := http.NewRequest("GET", url, nil)
	addAuthToken(authToken, req)

	res, err := c.HttpClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return -1
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var prList []pullRequestItem
	json.Unmarshal(body, &prList)
	return len(prList)
}

func addAuthToken(authToken string, req *http.Request) {
	if authToken != "" {
		req.Header.Add("Authorization", "token "+authToken)
	}
}

type repoListItem struct {
	Url  string
	Name string
}

type pullRequestItem struct {
	Id string
}

type repoItem struct {
	IssueCount     int    `json:"open_issues"`
	PullRequestUrl string `json:"pulls_url"`
	Name           string `json:"name" `
}
