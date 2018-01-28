package main

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type slackRequest struct {
	token       string
	teamID      string
	teamDomain  string
	channelID   string
	channelName string
	userID      string
	userName    string
	command     string
	text        string
	responseURL string
}

func (sr slackRequest) getUser() string {
	if sr.text != "" {
		return sr.text
	}
	return sr.userName
}

func (sr slackRequest) isValid() bool {
	return sr.token == os.Getenv("PR_VALID_TOKEN")
}

func getRequestFromSlashRequest(r *http.Request) slackRequest {
	bbody, _ := ioutil.ReadAll(r.Body)
	body := string(bbody)
	arr := strings.Split(body, "&")
	req := slackRequest{
		token:       getText(arr, "token"),
		teamID:      getText(arr, "team_id"),
		teamDomain:  getText(arr, "team_domain"),
		channelID:   getText(arr, "channel_id"),
		channelName: getText(arr, "channel_name"),
		userID:      getText(arr, "user_id"),
		userName:    getText(arr, "user_name"),
		command:     getText(arr, "command"),
		text:        getText(arr, "text"),
		responseURL: getText(arr, "response_url"),
	}
	return req
}

func getText(source []string, target string) string {
	for _, elem := range source {
		if strings.HasPrefix(elem, target) {
			arr := strings.Split(elem, "=")
			return arr[1]
		}
	}
	return ""
}
