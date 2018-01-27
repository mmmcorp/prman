package main

import "strings"

type pullRequests []pullRequest

type pullRequest struct {
	HTMLURL            string `json:"html_url"`
	Title              string `json:"title"`
	RequestedReviewers users  `json:"requested_reviewers"`
}

func (pr *pullRequest) isWIP() bool {
	return strings.HasPrefix(pr.Title, "WIP") || strings.HasPrefix(pr.Title, "(WIP)") || strings.HasPrefix(pr.Title, "[WIP]") || strings.HasPrefix(pr.Title, "【WIP】")
}

func (pr *pullRequest) isContainingAsReviewer(user string) bool {
	return pr.RequestedReviewers.contains(user)
}
