package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
)

type client struct {
	URL           *url.URL
	Org           string
	HTTPClient    *http.Client
	Authorization string
	Logger        *log.Logger
}

func decodeBody(resp *http.Response, out interface{}) error {
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	return decoder.Decode(out)
}

func newClient(org string, auth string) *client {
	client := new(client)
	u, _ := url.Parse("https://api.github.com")
	client.URL = u
	client.Org = org
	client.HTTPClient = &http.Client{}
	client.Authorization = auth
	client.Logger = log.New(os.Stdout, "logger: ", log.Lshortfile)
	return client
}

func (c *client) newRequest(method string, spath string, body io.Reader) (*http.Request, error) {
	u := *c.URL
	u.Path = path.Join(c.URL.Path, spath)
	req, err := http.NewRequest(method, u.String()+"?per_page=10000", body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", c.Authorization)
	return req, nil
}

func (c *client) getRepos() (*repos, error) {
	spath := fmt.Sprintf("/orgs/%v/repos", c.Org)
	req, _ := c.newRequest("GET", spath, nil)
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var repos repos
	if err := decodeBody(res, &repos); err != nil {
		return nil, err
	}
	return &repos, err
}

func (c *client) getIssueComments(repo string, number int) (*issueComments, error) {
	spath := fmt.Sprintf("/repos/%v/%v/issues/%v/comments", c.Org, repo, number)
	req, _ := c.newRequest("GET", spath, nil)
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var is issueComments
	if err := decodeBody(res, &is); err != nil {
		return nil, err
	}
	return &is, err
}

func (c *client) getPRs(repo string) (*pullRequests, error) {
	spath := fmt.Sprintf("/repos/%v/%v/pulls", c.Org, repo)
	req, _ := c.newRequest("GET", spath, nil)
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var prs pullRequests
	if err := decodeBody(res, &prs); err != nil {
		return nil, err
	}
	return &prs, err
}
