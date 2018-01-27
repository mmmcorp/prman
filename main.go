package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var c *client

func pullRequestController(w http.ResponseWriter, r *http.Request) {
	su, err := getUserFromSlashRequest(r)
	if err != nil {
		log.Println(err)
	}

	githubUser, err := convertSlackUserNameToGithubUserName(su)
	if err != nil {
		log.Println(err)
	}

	c = newClient(os.Getenv("PR_GITHUB_ORG"), os.Getenv("PR_GITHUB_TOKEN"))
	repos, err := c.getRepos()
	if err != nil {
		log.Println(err)
	}

	response, err := repos.getPullRequestsReviewRequestedFor(githubUser)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, fmt.Sprintf("{\"response_type\":\"in_channel\",\"text\":\"%v\"}", response))
}

func main() {
	http.HandleFunc("/", pullRequestController)
	http.ListenAndServe(":7000", nil)
}
