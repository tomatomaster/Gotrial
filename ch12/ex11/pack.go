// Copyright © 2017 Ryutarou Ono.

package main

import (
	"fmt"
	"net/url"
	"reflect"
)

func main() {
	var data struct {
		l   []string
		max int
		x   bool
	}
	data.l = []string{"test", "test", "123"}
	data.max = 10
	data.x = false

	url, _ := url.Parse("http://host")
	url = Pack(url, &data)
	fmt.Printf("%v", url)
}

func Pack(url *url.URL, ptr interface{}) *url.URL {
	v := reflect.ValueOf(ptr).Elem()
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i)
		val := v.Field(i)
		switch val.Kind() {
		case reflect.Array, reflect.Slice:
			for j := 0; j < val.Len(); j++ {
				setQuery(url, fieldInfo.Name, fmt.Sprintf("%v", val.Index(j)))
			}
		default:
			name := fieldInfo.Name
			setQuery(url, name, fmt.Sprintf("%v", val))
		}
	}
	return url
}

func setQuery(url *url.URL, q string, v string) *url.URL {
	query := url.Query()
	query.Add(q, v)
	url.RawQuery = query.Encode()
	return url
}
