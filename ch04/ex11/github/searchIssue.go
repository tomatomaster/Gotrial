package github

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

//searchURL is explained here https://developer.github.com/v3/search/#search-issues
const searchURL = Root + "/search/issues"

//SearchResultIssues reflects result of searching
type SearchResultIssues struct {
	TotalCount int `json:"toatal_counte"`
	Items      []*Issue
}

//Issue has some information related to issue.
type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // Markdown
}

//User is user informatioin related to issue.
type User struct {
	Login   string
	HTMLURL string `json:"created_at"`
}

//SearchIssues returns result of searching issue
func SearchIssues(terms []string) (*SearchResultIssues, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(searchURL + "?q=" + q)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}
	var result SearchResultIssues
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer resp.Body.Close()
	return &result, nil
}
