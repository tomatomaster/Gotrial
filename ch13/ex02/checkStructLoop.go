// Copyright © 2017 Ryutarou Ono.

package main

import (
	"log"
	"reflect"
	"unsafe"
)

func IsLoop(x interface{}) bool {
	seen := make(map[comparison]bool)
	return loop(reflect.ValueOf(x), seen)
}

type comparison struct {
	x unsafe.Pointer
}

func loop(x reflect.Value, seen map[comparison]bool) bool {
	//log.Print(x.Kind()) For debug
	switch x.Kind() {
	case reflect.Ptr, reflect.Interface:
		return loop(x.Elem(), seen)

	case reflect.Struct:
		for i, n := 0, x.NumField(); i < n; i++ {
			if loop(x.Field(i), seen) {
				return true
			}
		}

	default:
		if x.CanAddr() {
			xptr := unsafe.Pointer(x.UnsafeAddr())
			c := comparison{xptr} //d, e, link
			if seen[c] {
				log.Printf("%v has been allready Seen", x)
				return true
			}
			seen[c] = true
		}
	}
	return false
}
