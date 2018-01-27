package main

import (
	"fmt"
	"sync"
)

type repos []repo

type repo struct {
	Name string `json:"name"`
}

func (repos repos) getPullRequestsReviewRequestedFor(user string) (string, error) {
	response := "WIPでない、レビュアーに指定されているPRをお知らせします。\n"
	var wg sync.WaitGroup
	for _, r := range repos {
		wg.Add(1)
		go func(r repo) {
			res, err := r.getPullRequestReviewRequestedFor(user)
			if err != nil {
				panic(err)
			}
			response += res
			wg.Done()
		}(r)
	}
	wg.Wait()
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
			response += fmt.Sprintf("* %v : %v(%v)\n", repo.Name, pr.Title, pr.URL)
		}
	}
	return response, nil
}
