package config

import (
	"bufio"
	"go/build"
	"io/ioutil"
	"os"
	"strings"
)

func SetAccessToken(token string) {
	var currentLines = getAllLines()

	if len(currentLines) == 1 {
		currentLines = append(currentLines, token)
	} else if len(currentLines) == 2 {
		currentLines[1] = token
	}

	gopath := getGoPath()
	var newLines = strings.Join(currentLines, "\n")
	if token != "" {
		ioutil.WriteFile(gopath+"/github_summary.data", []byte(newLines), 0777)
	}
}

func SetUserName(userName string) {
	gopath := getGoPath()
	var currentLines = getAllLines()
	if len(currentLines) == 0 {
		ioutil.WriteFile(gopath+"/github_summary.data", []byte(userName), 0777)
		return
	}

	currentLines[0] = userName
	var newLines = strings.Join(currentLines, "\n")
	ioutil.WriteFile(gopath+"/github_summary.data", []byte(newLines), 0777)
}

func getGoPath() string {
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		gopath = build.Default.GOPATH
	}
	return gopath
}

func GetUserNameAndAuthToken() (string, string) {
	lines := getAllLines()
	if len(lines) == 1 {
		return lines[0], ""
	} else if len(lines) == 2 {
		return lines[0], lines[1]
	}

	return "", ""
}

func getAllLines() []string {
	gopath := getGoPath()
	file, err := os.Open(gopath + "/github_summary.data")
	if err != nil {
		return nil
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
