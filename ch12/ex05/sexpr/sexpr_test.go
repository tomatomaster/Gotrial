// Copyright © 2017 Ryutarou Ono.

package sexpr

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

type Test struct {
	String1 string
	Int1    int
	Bool1   bool
	Map1    map[string]string
	Array1  []string
}

func TestMarshal(t *testing.T) {
	sample := Test{
		String1: "test",
		Int1:    1,
		Bool1:   false,
		Map1: map[string]string{
			"testkey": "testval",
		},
		Array1: []string{"test"},
	}
	marshaled, err := Marshal(sample)
	if err != nil {
		log.Fatal(err)
	}
	marshaledString := fmt.Sprintf("%s", string(marshaled))
	fmt.Print(marshaledString)
	var unmarshaled Test
	json.Unmarshal(marshaled, &unmarshaled)

	if fmt.Sprintf("%v", unmarshaled) != fmt.Sprintf("%v", sample) {
		t.Errorf("may not marshal type of struct to type of json")
	}
}
