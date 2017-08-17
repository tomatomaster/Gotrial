// Copyright © 2017 Ryutarou Ono.

package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"sort"
)

const templ = `<table>
<tr style='text-align: left'>
<th><a href="?sort=title">Title</a></th>
<th><a href="?sort=artist">Artist</a></th>
<th><a href="?sort=album">Album</a></th>
<th>Year</th>
<th>Length</th>
</tr>
{{range .}}<tr>
<td>{{.Title}}</td>
<td>{{.Artist}}</td>
<td>{{.Album}}</td>
<td>{{.Year}}</td>
<td>{{.Length}}</td>
</tr>
{{end}}`

var report = template.Must(template.New("trackList").Parse(templ))

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("sort")
	switch query {
	case "title":
		sort.Sort(byTitle(tracks))
	case "artist":
		sort.Sort(byArtist(tracks))
	case "album":
		sort.Sort(byAlbum(tracks))
	default:
		fmt.Println("query has not defined yet or nil")
	}
	if err := report.Execute(w, tracks); err != nil {
		log.Fatal(err)
	}
}
