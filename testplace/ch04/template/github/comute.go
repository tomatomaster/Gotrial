package github

import (
	"bytes"
	"log"
	"net/http"
	"time"
)

func postJSON(u string, json []byte, auth AuthInfo) (*http.Response, error) {
	req, err := http.NewRequest("POST", u, bytes.NewReader(json))
	if err != nil {
		log.Fatal(err)
		return &http.Response{}, err
	}
	req.Header.Add("Accept", "application/vnd.github.v3.text-match+json")
	req.SetBasicAuth(auth.Name, auth.Pass)
	client := &http.Client{Timeout: time.Duration(10) * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	return resp, nil
}
