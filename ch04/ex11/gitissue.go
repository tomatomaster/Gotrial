package main

import "./github"
import "fmt"

func main() {
	request := github.IssueRequest{Title: "TestTitle", Body: "TestBody", Milestone: 1}
	resp, err := github.CreateIssues("tomatomaster", "githubapitest", &request)
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Printf("resp %v", resp)
	}
}
