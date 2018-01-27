package main

import (
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

func getUserFromSlashRequest(r *http.Request) (string, error) {
	bbody, _ := ioutil.ReadAll(r.Body)
	body := string(bbody)
	arr := strings.Split(body, "&")
	if getText(arr, "text") != "" {
		return getText(arr, "text"), nil
	}
	if getText(arr, "user_name") != "" {
		return getText(arr, "user_name"), nil
	}
	return "", errors.New("user_nameパラメータがありません。")
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
