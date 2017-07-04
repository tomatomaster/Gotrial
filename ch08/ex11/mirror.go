// Copyright © 2017 Ryutarou Ono.

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	urls := os.Args
	fmt.Printf("Response: %s\n", fetch(urls...))
}

var done = make(chan struct{})

func fetch(urls ...string) string {
	response := make(chan string)
	for _, url := range urls {
		go func() {
			select {
			case <-done:
				return
			default:
				response <- request(url)
			}
		}()
	}
	return <-response
}

func request(hostname string) (response string) {
	request, err := http.NewRequest("Get", hostname, nil)
	if err != nil {
		log.Fatal(err)
	}
	cancel := make(chan struct{})
	request.Cancel = cancel
	go func(cancel chan struct{}, request *http.Request) {
		for {
			select {
			case <-done:
				fmt.Printf("Close %s \n", request.URL)
				close(cancel)
			}
		}
	}(cancel, request)
	if isCanceled() {
		return
	}
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Printf("Request has already closed")
		return
	}
	done <- struct{}{}
	return resp.Request.URL.String()
}

func isCanceled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}
