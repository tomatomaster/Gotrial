// Copyright © 2017 Ryutarou Ono.

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"sync"
)

type depths struct {
	Import string   `json:ImportPath`
	Depth  []string `json:"Deps"`
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) //Use all cpu core
	wantCheck := os.Args[1]
	out, err := exec.Command("go", "list", "...").Output()
	if err != nil {
		log.Fatal(err)
	}
	packages := strings.Split(fmt.Sprintf("%s", out), "\n")
	var wg sync.WaitGroup
	for _, packageName := range packages {
		wg.Add(1)
		go func(packageName string) {
			defer wg.Done()
			out, err := exec.Command("go", "list", "-json", packageName).Output()
			if err != nil {
				log.Fatal(err)
			}
			data := new(depths)
			json.Unmarshal(out, data)
			//All Package Dependency
			for _, dependency := range data.Depth {
				if dependency == wantCheck {
					fmt.Printf("%s\n", packageName)
				}
			}
		}(packageName)
	}
	wg.Wait()
}
