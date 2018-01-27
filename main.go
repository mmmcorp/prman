package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

var c *client

func pullRequestController(w http.ResponseWriter, r *http.Request) {
	// rはこの関数が終わるとfreeされるので先にreqオブジェクトだけ作る
	req := getRequestFromSlashRequest(r)
	firstResponse(w)
	// タイムアウトを避けるために先にfirstResponseだけ返す
	// この関数が終わらないとSlackで表示されないので、並列実行する
	go secondResponse(req)
}

func firstResponse(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "{\"response_type\":\"in_channel\",\"text\":\"PRを調べています..\"}")
}

func secondResponse(req slackRequest) {
	githubUser, err := convertSlackUserNameToGithubUserName(req.getUser())
	if err != nil {
		response(req.responseURL, err.Error())
	}

	c = newClient(os.Getenv("PR_GITHUB_ORG"), os.Getenv("PR_GITHUB_TOKEN"))

	repos, err := c.getRepos()
	if err != nil {
		response(req.responseURL, err.Error())
	}

	res, err := repos.getPullRequestsReviewRequestedFor(githubUser)
	if err != nil {
		response(req.responseURL, err.Error())
	}

	response(req.responseURL, res)
}

func response(u string, body string) {
	b := fmt.Sprintf("{\"response_type\":\"in_channel\",\"text\":\"%v\"}", body)
	decodedURL, _ := url.QueryUnescape(u)
	req, err := http.NewRequest("POST", decodedURL, bytes.NewBuffer([]byte(b)))
	if err != nil {
		panic(err)
	}
	cl := &http.Client{}
	cl.Do(req)
}

func main() {
	http.HandleFunc("/", pullRequestController)
	http.ListenAndServe(":7000", nil)
}
