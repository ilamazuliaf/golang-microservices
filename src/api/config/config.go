package config

import "os"

const (
	apiGithubAccessToken = "SECRET_GITHUB_ACCESS_TOKEN"
)

var (
	// Read token github from system env
	// use set in windows and export in linux
	githubAccessToken = os.Getenv(apiGithubAccessToken)
)

func GetGithubAccessToken() string {
	return  githubAccessToken
}