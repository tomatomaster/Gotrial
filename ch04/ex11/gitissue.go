package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"./github"
)

func main() {
	var mode string
	var owner, repo string
	flag.StringVar(&mode, "mode", "", "")
	flag.StringVar(&owner, "owner", "", "repository owner")
	flag.StringVar(&repo, "repo", "", "repository name")
	flag.Parse()

	imap := make(map[string]string)
	if mode == "create" {
		inputs := inputCommon(imap)
		request := github.IssueFields{Title: inputs["title"], Body: inputs["body"]}
		auth := github.AuthInfo{Name: inputs["name"], Pass: inputs["pass"]}
		resp, err := github.CreateIssues(owner, repo, &request, auth)
		if err != nil {
			fmt.Print(err)
		} else {
			fmt.Printf("resp %v", resp.Body)
		}
	} else if mode == "edit" {
		inputs := inputEditIssue(imap)
		request := github.IssueFields{Title: inputs["title"], Body: inputs["body"]}
		auth := github.AuthInfo{Name: inputs["name"], Pass: inputs["pass"]}
		resp, err := github.EditIssue(owner, repo, inputs["no"], &request, auth)
		if err != nil {
			fmt.Print(err)
		} else {
			fmt.Printf("resp %v", resp.Body)
		}
	} else if mode == "close" {
		inputs := inputCloseIssue(imap)
		auth := github.AuthInfo{Name: inputs["name"], Pass: inputs["pass"]}
		resp, err := github.CloseIssue(owner, repo, inputs["no"], auth)
		if err != nil {
			fmt.Print(err)
		} else {
			fmt.Printf("resp %v", resp.Body)
		}
	} else if mode == "search" {
		result, err := github.SearchIssues([]string{"test", "http"})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%v\n", result.TotalCount)
	}
}

func inputCommon(inputs map[string]string) map[string]string {
	fmt.Println("Enter the issue title")
	inputs["title"] = nextLine()
	fmt.Println("Enter the issue body")
	inputs["body"] = nextLine()
	return inputs
}

func inputAuthInfo(inputs map[string]string) map[string]string {
	fmt.Println("Enter username:")
	inputs["name"] = nextLine()
	fmt.Println("Enter password:")
	inputs["pass"] = nextLine()
	return inputs
}

func inputEditIssue(inputs map[string]string) map[string]string {
	fmt.Println("Enter the issue no")
	inputs["no"] = nextLine()
	inputs = inputCommon(inputs)
	inputs = inputAuthInfo(inputs)
	return inputs
}

func inputCloseIssue(inputs map[string]string) map[string]string {
	fmt.Println("Enter the issue no")
	inputs["no"] = nextLine()
	inputs = inputAuthInfo(inputs)
	return inputs
}

func nextLine() string {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	return sc.Text()
}
