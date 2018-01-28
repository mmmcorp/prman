package main

import "strings"

type pullRequests []pullRequest

type pullRequest struct {
	HTMLURL string `json:"html_url"`
	Title   string `json:"title"`
	Number  int    `json:"number"`
	Body    string `json:"body"`
}

func (pr *pullRequest) isWIP() bool {
	return strings.HasPrefix(pr.Title, "WIP") || strings.HasPrefix(pr.Title, "(WIP)") || strings.HasPrefix(pr.Title, "[WIP]") || strings.HasPrefix(pr.Title, "【WIP】")
}

func (pr *pullRequest) mustBeReviewedBy(user string) bool {
	return strings.Contains(pr.Body, user)
}
