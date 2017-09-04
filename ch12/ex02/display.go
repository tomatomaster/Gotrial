// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 333.

// Package display provides a means to display structured data.
package main

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"strconv"
)

//!+Display

func Display(name string, x interface{}) {
	fmt.Printf("Display %s (%T):\n", name, x)
	display := new(newDisplay)
	display.display(name, reflect.ValueOf(x))
}

//!-Display

// formatAtom formats a value without inspecting its internal structure.
// It is a copy of the the function in gopl.io/ch11/format.
func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
	// ...floating-point and complex cases omitted for brevity...
	case reflect.Bool:
		if v.Bool() {
			return "true"
		}
		return "false"
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr,
		reflect.Slice, reflect.Map:
		return v.Type().String() + " 0x" +
			strconv.FormatUint(uint64(v.Pointer()), 16)
	case reflect.Struct:
		var fieldPath string
		fieldPath = "\n"
		for i := 0; i < v.NumField(); i++ {
			fieldPath += fmt.Sprintf("\t%s.%s=%v(%s)\n",
				v.Type(), v.Field(i).String(), formatAtom(v.Field(i)), v.Field(i).Type().String())
		}
		return fieldPath
	case reflect.Array:
		var arrayVal string
		for i := 0; i < v.Len(); i++ {
			arrayVal += fmt.Sprintf("[%d] %v", i, v.Index(i))
		}
		return arrayVal
	default: // reflect.Array, reflect.Struct, reflect.Interface
		return v.Type().String() + " value"
	}
}

type newDisplay struct {
	limit int8
}

//!+display
func (n newDisplay) display(path string, v reflect.Value) error {
	if n.limit > math.MaxInt8-1 {
		return errors.New("This value has cyclic.\n")
	}
	n.limit++

	switch v.Kind() {
	case reflect.Invalid:
		fmt.Printf("%s = invalid\n", path)
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			err := n.display(fmt.Sprintf("%s[%d]", path, i), v.Index(i))
			if err != nil {
				fmt.Printf("%v", err)
				return nil
			}
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			fieldPath := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)
			err := n.display(fieldPath, v.Field(i))
			if err != nil {
				fmt.Printf("%v", err)
				return nil
			}
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			err := n.display(fmt.Sprintf("%s[%s]", path,
				formatAtom(key)), v.MapIndex(key))
			if err != nil {
				fmt.Printf("%v", err)
				return nil
			}
		}
	case reflect.Ptr:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			err := n.display(fmt.Sprintf("(*%s)", path), v.Elem())
			if err != nil {
				fmt.Printf("%v", err)
				return nil
			}
		}
	case reflect.Interface:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			fmt.Printf("%s.type = %s\n", path, v.Elem().Type())
			err := n.display(path+".value", v.Elem())
			if err != nil {
				fmt.Printf("%v", err)
				return nil
			}
		}
	default: // basic types, channels, funcs
		fmt.Printf("%s = %s\n", path, formatAtom(v))
	}
	return nil
}

//!-display
