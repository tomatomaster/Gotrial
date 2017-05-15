package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

//IssuesURL is explained here https://developer.github.com/v3/search/#search-issues
const IssuesURL = "https://api.github.com/search/issues"

//IssuesSearchResult reflects result of searching
type IssuesSearchResult struct {
	TotalCount int `json:"toatal_counte"`
	Items      []*Issue
}

//Issue is
type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // Markdown
}

//User is
type User struct {
	Login   string
	HTMLURL string `json:"created_at"`
}

//SearchIssues　returns result of searching issue
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}
	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

//
type IssueRequest struct {
	Title     string   `json:"title,omitempty"`
	Body      string   `json:"body,omitempty"`
	Milestone int      `json:"milestone,omitempty"`
	Assignees []string `json:"assignees,omitempty"`
}

func CreateIssues(owner, repo string, issue *IssueRequest) (*http.Request, error) {
	u := fmt.Sprintf("repos/%v/%v/issues", owner, repo)
	url := "https://api.github.com/" + u
	log.Printf("Access to %v", url)
	b, _ := json.Marshal(issue)
	body := bytes.NewReader(b)
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return req, nil
	}
	return nil, err
}
