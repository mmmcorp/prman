package main

import (
	"fmt"
	"strings"
	"sync"
)

type repos []repo

type repo struct {
	Name string `json:"name"`
}

func (repos repos) getPullRequestsReviewRequestedFor(user string) (string, error) {
	response := "WIPでない、レビュアーに指定されているPRをお知らせします。\n"
	response += "```\n"
	var wg sync.WaitGroup
	var updated bool
	for _, r := range repos {
		wg.Add(1)
		go func(r repo) {
			res, err := r.getPullRequestReviewRequestedFor(user)
			if err != nil {
				panic(err)
			}
			if res != "" {
				response += res
				updated = true
			}
			wg.Done()
		}(r)
	}
	wg.Wait()
	response += "```\n"
	if !updated {
		response = strings.Replace(response, "`", "", -1)
		response += "レビュー待ちになっているPRはありませんでした!! :tada:"
	}
	return response, nil
}

func (repo repo) getPullRequestReviewRequestedFor(user string) (string, error) {
	response := ""
	prs, err := c.getPRs(repo.Name)
	if err != nil {
		return "", err
	}
	for _, pr := range *prs {
		if !pr.isWIP() && pr.isContainingAsReviewer(user) {
			response += fmt.Sprintf("* %v : %v(%v)\n", repo.Name, pr.Title, pr.HTMLURL)
		}
	}
	return response, nil
}
