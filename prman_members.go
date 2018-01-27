package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"strings"
)

type prmanMembers struct {
	Members []string `json:"members"`
}

func (pr *prmanMembers) getUserNameFrom(slUserName string) (string, error) {
	for _, member := range pr.Members {
		arr := strings.Split(member, ":")
		sl := arr[0]
		gh := arr[1]
		if sl == slUserName {
			return gh, nil
		}
	}
	return "", errors.New("Userが見つかりませんでした。")
}

func convertSlackUserNameToGithubUserName(slUserName string) (string, error) {
	file, err := ioutil.ReadFile("./prman-members.json")
	if err != nil {
		return "", errors.New("prman-members.jsonがありません。")
	}
	var prmanMembers prmanMembers
	json.Unmarshal(file, &prmanMembers)
	return prmanMembers.getUserNameFrom(slUserName)
}
