package main

import (
	"strings"
)

type issueComments []issueComment

type issueComment struct {
	User    user   `json:"user"`
	HTMLURL string `json:"html_url"`
	Body    string `json:"body"`
}

func (ics issueComments) selectFrom(pr pullRequest) issueComments {
	var res []issueComment
	for _, comment := range ics {
		if strings.Contains(comment.HTMLURL, pr.HTMLURL) {
			res = append(res, comment)
		}
	}
	return res
}

func (ics issueComments) hasReviewCommentFrom(username string) bool {
	for _, comment := range ics {
		if comment.User.Login == username && strings.Contains(comment.Body, "ロジック含めて詳細にコードをレビューした") {
			return true
		}
	}
	return false
}
