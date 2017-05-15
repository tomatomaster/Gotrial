package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type jResult struct {
	Poster string
}

const baseURL = "http://www.omdbapi.com"

func main() {
	query := strings.Join(os.Args[1:], "+")
	url := fmt.Sprintf("%s/?t=%s", baseURL, query)
	fmt.Println(url)
	download(getImgURL(url))
}

func getImgURL(url string) string {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		log.Printf("Faile to access %v", response.StatusCode)
	}
	var result jResult
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		log.Fatal(err)
	}
	return result.Poster
}

func download(url string) {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		log.Printf("Faile to access %v", response.StatusCode)
	}
	file, err := os.Create("poster.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	io.Copy(file, response.Body)
}
