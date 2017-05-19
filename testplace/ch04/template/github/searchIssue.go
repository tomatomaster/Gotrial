package github

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

//searchURL is explained here https://developer.github.com/v3/search/#search-issues
const searchURL = Root + "/search/issues"

//SearchResultIssues reflects result of searching
type SearchResultIssues struct {
	TotalCount int `json:"toatal_counte"`
	Items      []*Issue
}

type RepositoryIssues struct {
	ID int
}

//Issue has some information related to issue.
type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	Milestone *Milestone
	CreatedAt time.Time `json:"created_at"`
	Body      string    // Markdown
}

//User is user informatioin related to issue.
type User struct {
	Login   string
	HTMLURL string `json:"created_at"`
}

//Milestone is milestone information related to issue.
type Milestone struct {
	URL   string
	State string
	ID    int
	Title string `json:"title"`
}

//SearchIssues returns result of searching issue
func SearchIssues(owner, repo string) (*RepositoryIssues, error) {
	url := fmt.Sprintf("/repos/%v/%v/issues", owner, repo)
	resp, err := http.Get(Root + url)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}
	var result RepositoryIssues
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer resp.Body.Close()
	return &result, nil
}
