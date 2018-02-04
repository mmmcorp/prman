package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"
)

var c *client

func pullRequestController(w http.ResponseWriter, r *http.Request) {
	c = newClient(os.Getenv("PR_GITHUB_ORG"), os.Getenv("PR_GITHUB_TOKEN"))

	// rはこの関数が終わるとfreeされるので先にreqオブジェクトだけ作る
	req := getRequestFromSlashRequest(r)

	c.Logger.Printf("アクセスを受け付けました。requested by : %v request: %v+\n", req.getUser(), req)
	firstResponse(w)
	c.Logger.Printf("firstResponseを返しました。requested by: %v \n", req.getUser())

	// タイムアウトを避けるために先にfirstResponseだけ返す
	// この関数が終わらないとSlackで表示されないので、並列実行する
	c.Logger.Printf("secondResponse関数を開始します。requested by : %v \n", req.getUser())
	go secondResponse(req)
}

func firstResponse(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "{\"response_type\":\"in_channel\",\"text\":\"PRを調べています..\"}")
}

// 本体
func secondResponse(req slackRequest) {
	start := time.Now()
	c.Logger.Printf("isValid関数を開始します。requested by : %v \n", req.getUser())
	if !req.isValid() {
		c.Logger.Printf("Errorが起こりました。err: %v, requested by: %v \n", "invalid token", req.getUser())
		response(req.responseURL, "invalid token")
		return
	}

	c.Logger.Printf("convertSlackUserNameToGithubUserName関数を開始します。requested by : %v \n", req.getUser())
	githubUser, err := convertSlackUserNameToGithubUserName(req.getUser())
	if err != nil {
		c.Logger.Printf("Errorが起こりました。err: %v, requested by: %v \n", err.Error(), req.getUser())
		response(req.responseURL, err.Error())
		return
	}
	c.Logger.Printf("githubUser: %v , requested by : %v \n", githubUser, req.getUser())

	c.Logger.Printf("getRepos関数を開始します。requested by : %v \n", req.getUser())
	repos, err := c.getRepos()
	if err != nil {
		c.Logger.Printf("Errorが起こりました。err: %v, requested by: %v \n", err.Error(), req.getUser())
		response(req.responseURL, err.Error())
		return
	}
	c.Logger.Printf("repos: %v , requested by : %v \n", repos, req.getUser())

	c.Logger.Printf("getPullRequestsReviewRequestedFor関数を開始します。requested by : %v \n", req.getUser())
	res, err := repos.getPullRequestsReviewRequestedFor(githubUser)
	if err != nil {
		c.Logger.Printf("Errorが起こりました。err: %v, requested by: %v \n", err.Error(), req.getUser())
		response(req.responseURL, err.Error())
		return
	}
	c.Logger.Printf("res: %v , requested by : %v \n", res, req.getUser())
	end := time.Now()
	if req.isDebug() {
		res += fmt.Sprintf("start: %v\n", start)
		res += fmt.Sprintf("end: %v\n", end)
	}

	c.Logger.Printf("response関数を開始します。requested by : %v \n", req.getUser())
	response(req.responseURL, res)
	c.Logger.Printf("responseを返しました。requested by : %v \n", req.getUser())
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
