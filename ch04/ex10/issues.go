package main

import (
	"fmt"
	"log"

	"os"

	"time"

	"./github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, t := range classifyByDate(result.Items) {
		if t.issue == nil {
			continue
		}
		issue := t.issue
		fmt.Printf("%s(%v)\n#%-5d %9.9s %.55s \n\n", t.date, issue.CreatedAt, issue.Number, issue.User.Login, issue.Title)
	}
}

type eIssue struct {
	issue *github.Issue
	date  string
}

func classifyByDate(issues []*github.Issue) []eIssue {
	result := make([]eIssue, len(issues))
	nYear, nMonth, _ := time.Now().Date()
	for _, issue := range issues {
		year, month, _ := issue.CreatedAt.Date()
		if nYear-year > 0 {
			result = append(result, eIssue{issue, "More Than Year"})
			continue
		}
		if nMonth-month > 0 {
			result = append(result, eIssue{issue, "Less Than Year"})
			continue
		}
		result = append(result, eIssue{issue, "Less Than Month"})
	}
	return result
}
