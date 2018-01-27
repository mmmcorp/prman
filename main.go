package main

import (
	"fmt"
	"net/http"
	"os"
)

var c *client

func pullRequestController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	su, err := getUserFromSlashRequest(r)
	if err != nil {
		response(w, err.Error())
	}

	githubUser, err := convertSlackUserNameToGithubUserName(su)
	if err != nil {
		response(w, err.Error())
	}

	c = newClient(os.Getenv("PR_GITHUB_ORG"), os.Getenv("PR_GITHUB_TOKEN"))
	repos, err := c.getRepos()
	if err != nil {
		response(w, err.Error())
	}

	res, err := repos.getPullRequestsReviewRequestedFor(githubUser)
	if err != nil {
		response(w, err.Error())
	}

	response(w, res)
}

func response(w http.ResponseWriter, body string) {
	fmt.Fprintf(w, fmt.Sprintf("{\"response_type\":\"in_channel\",\"text\":\"%v\"}", body))
}

func main() {
	http.HandleFunc("/", pullRequestController)
	http.ListenAndServe(":7000", nil)
}
