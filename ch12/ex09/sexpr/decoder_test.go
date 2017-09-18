// Copyright © 2017 Ryutarou Ono.

package sexpr

import (
	"testing"

	"bytes"
)

type Data struct {
	Title, SubTitle string
	Year            int
}

func TestDecoder_UnmarshalReader(t *testing.T) {
	samples := []struct {
		Input string
		Want  Data
	}{
		{
			Input: `((Title "test") (SubTitle "Fxxk you") (Year 123))`,
			Want: Data{
				Title:    "test",
				SubTitle: "Fxxk you",
				Year:     123,
			}},
		{
			Input: `((Title "test") (SubTitle "") (Year))`,
			Want: Data{
				Title:    "test",
				SubTitle: "",
				Year:     0,
			}},
	}
	for _, s := range samples {
		var data Data
		reader := bytes.NewReader([]byte(s.Input))
		decoder := NewReader(reader)
		decoder.UnmarshalReader(&data)
		if data.Title != s.Want.Title || data.SubTitle != s.Want.SubTitle || data.Year != s.Want.Year {
			t.Fatalf("Error %s %s %d", data.Title, data.SubTitle, data.Year)
		}
	}
}
