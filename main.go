package main

import (
	"bytes"
	"fmt"
	"net/http"
)

// PullRequestServer serves pull requests
func PullRequestServer(w http.ResponseWriter, r *http.Request) {
	bufbody := new(bytes.Buffer)
	bufbody.ReadFrom(r.Body)
	body := bufbody.String()
	fmt.Fprintf(w, body)
	fmt.Fprintf(w, "Hello, World\n")
}
func main() {
	http.HandleFunc("/", PullRequestServer)
	http.ListenAndServe(":7000", nil)
}
