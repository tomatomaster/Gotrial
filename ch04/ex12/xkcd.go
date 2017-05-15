package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

var indexes map[int]xkcdResult

func main() {
	indexes = make(map[int]xkcdResult)
	fmt.Println("Enter index no.")
	var sc = bufio.NewScanner(os.Stdin)
	var index int
	var err error
	for sc.Scan() {
		i := sc.Text()
		if i != "" {
			if index, err = strconv.Atoi(sc.Text()); err != nil {
				log.Fatal(err)
			}
			showResult(index)
		}
	}
}

func showResult(index int) {
	result, err := get(index)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v\n\n", result)
}

type xkcdResult struct {
	Month      string
	Num        int
	Link       string
	Year       string
	News       string
	SafeTitle  string `json:"safe_title"`
	Transcript string
	Alt        string
	Img        string
	Title      string
	Day        string
}

func get(no int) (xkcdResult, error) {
	content := indexes[no]
	if content.Title != "" { //既にxkcdから情報を取得していた場合
		return content, nil
	}
	u := fmt.Sprintf("https://xkcd.com/%d/info.0.json", no)
	resp, err := http.Get(u)
	if err != nil {
		log.Fatal(err)
		return content, nil
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return content, nil
	}

	var result xkcdResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		log.Fatal(err)
		resp.Body.Close()
		return content, nil
	}
	resp.Body.Close()
	indexes[no] = result
	return result, nil
}
