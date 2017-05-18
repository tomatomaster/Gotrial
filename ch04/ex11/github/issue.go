package github

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//IssueFields has some fileds relative to issue.
type IssueFields struct {
	Title     string   `json:"title,omitempty"`
	Body      string   `json:"body,omitempty"`
	Milestone int      `json:"milestone,omitempty"`
	Assignees []string `json:"assignees,omitempty"`
	Labels    []string `json:"labels,omitempty"`
	State     string   `json:"state,omitempty"`
}

//CreateIssues create issue
func CreateIssues(owner, repo string, issue *IssueFields, auth AuthInfo) (*http.Response, error) {
	u := fmt.Sprintf("/repos/%v/%v/issues", owner, repo)
	url := Root + u
	log.Printf("Access to %v", url)
	b, _ := json.Marshal(issue)
	request, err := postJSON(url, b, auth)
	return request, err
}

//EditIssue edit issue
func EditIssue(owner, repo, number string, issue *IssueFields, auth AuthInfo) (*http.Response, error) {
	u := fmt.Sprintf("/repos/%v/%v/issues/%v", owner, repo, number)
	url := Root + u
	b, _ := json.Marshal(issue)
	request, err := postJSON(url, b, auth)
	return request, err
}

//CloseIssue close issue
func CloseIssue(owner, repo, number string, auth AuthInfo) (*http.Response, error) {
	issue := IssueFields{State: "closed"}
	return EditIssue(owner, repo, number, &issue, auth)
}
