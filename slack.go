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
	for _, elem := range arr {
		if strings.HasPrefix(elem, "user_name") {
			// elem will be user_name=yagi
			return strings.Split(elem, "=")[1], nil
		}
	}
	return "", errors.New("user_nameパラメータがありません。")
}
