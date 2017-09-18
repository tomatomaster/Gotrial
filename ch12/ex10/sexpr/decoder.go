// Copyright © 2017 Ryutarou Ono.

package sexpr

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"reflect"
	"strconv"
	"strings"
	"text/scanner"
)

type lexer struct {
	scan  scanner.Scanner
	token rune
}

func (lex *lexer) next()        { lex.token = lex.scan.Scan() }
func (lex *lexer) text() string { return lex.scan.TokenText() }

func (lex *lexer) consume(want rune) {
	if lex.token != want {
		panic(fmt.Sprintf("got %q, want %q", lex.text(), want))
	}
	lex.next()
}

type Decoder struct {
	r io.Reader
}

func NewReader(reader io.Reader) *Decoder {
	return &Decoder{reader}
}

func (d Decoder) UnmarshalReader(out interface{}) (err error) {
	lex := &lexer{scan: scanner.Scanner{Mode: scanner.GoTokens}}
	lex.scan.Init(d.r)
	lex.next()
	defer func() {
		if x := recover(); x != nil {
			err = fmt.Errorf("error at %s: %v", lex.scan.Position, x)
		}
	}()
	read(lex, reflect.ValueOf(out).Elem())
	return nil
}

func Unmarshal(data []byte, out interface{}) (err error) {
	lex := &lexer{scan: scanner.Scanner{Mode: scanner.GoTokens}}
	lex.scan.Init(bytes.NewReader(data))
	lex.next()
	defer func() {
		if x := recover(); x != nil {
			err = fmt.Errorf("error at %s: %v", lex.scan.Position, x)
		}
	}()
	read(lex, reflect.ValueOf(out).Elem())
	return nil
}

/**
((Text:"text") (Year:12))
*/
func read(lex *lexer, v reflect.Value) {
	log.Printf("%s %v", lex.text(), lex.token)
	switch lex.token {
	case scanner.Ident: //Text, Year or nil
		if lex.text() == "nil" {
			v.Set(reflect.Zero(v.Type()))
			lex.next()
			return
		} else if lex.text() == "t" {
			v.SetBool(true)
			lex.next()
			return
		} else {

		}
	case scanner.String:
		s, _ := strconv.Unquote(lex.text()) //"text" -> text
		v.SetString(s)
		lex.next()
		return
	case scanner.Int:
		i, _ := strconv.Atoi(lex.text())
		v.SetInt(int64(i))
		lex.next()
		return
	case scanner.Float:
		i, _ := strconv.ParseFloat(lex.text(), 64)
		v.SetFloat(float64(i))
		lex.next()
		return

	case '(':
		lex.next() // consume '('
		readList(lex, v)
		lex.next() // consume ')'
		return

	case '#':
		readComplex(lex, v)
		return
	}
	panic(fmt.Sprintf("unexpected token %q", lex.text()))
}

func readComplex(lex *lexer, v reflect.Value) {
	lex.next()
	// consume '#'
	lex.next()
	// consume 'C'
	lex.next()
	// '('
	r, _ := strconv.ParseFloat(lex.text(), 128)
	lex.next()
	// consume real
	i, _ := strconv.ParseFloat(lex.text(), 128)
	lex.next()
	// consume image
	lex.next()
	// consume ')'
	c := complex128(complex(r, i))
	v.SetComplex(c)
}

/**

 */
func readList(lex *lexer, v reflect.Value) {
	switch v.Kind() {
	case reflect.Struct:
		for !endList(lex) {
			lex.consume('(')
			if lex.token != scanner.Ident {
				panic(fmt.Sprintf("got token %q, want filed name", lex.text()))
			}
			name := lex.text()
			lex.next()
			read(lex, v.FieldByName(name)) // struct
			lex.consume(')')
		}

	case reflect.Array:
		for i := 0; !endList(lex); i++ {
			read(lex, v.Index(i))
		}

	case reflect.Slice:
		for !endList(lex) {
			item := reflect.New(v.Type().Elem()).Elem()
			read(lex, item)
			v.Set(reflect.Append(v, item))
		}

	case reflect.Map:
		v.Set(reflect.MakeMap(v.Type())) //make map
		for !endList(lex) {
			lex.consume('(')
			key := reflect.New(v.Type().Key()).Elem()
			read(lex, key)
			value := reflect.New(v.Type().Elem()).Elem()
			read(lex, value)
			v.SetMapIndex(key, value)
			lex.consume(')')
		}

	case reflect.Interface:
		t, _ := strconv.Unquote(lex.text()) //Information
		log.Printf("Information %s ", t)
		value := reflect.New(getSupportedType(t)).Elem()
		lex.next()
		read(lex, value)
		v.Set(value)

	default:
		panic(fmt.Sprintf("cannot decode list into %v", v.Type()))
	}
}

func endList(lex *lexer) bool {
	switch lex.token {
	case scanner.EOF:
		panic("end of file")
	case ')':
		return true
	}
	return false
}

/**
[]int
[4]int
map[string][int]
*/
func getSupportedType(t string) reflect.Type {
	log.Printf("Get type %s", t)
	ty, e := getBasicTypea(t)
	if e == nil {
		return ty
	}

	if strings.HasPrefix(t, "[]") {
		ty, _ = getBasicTypea(t[2:])
		return reflect.SliceOf(ty)
	}

	if strings.HasPrefix(t, "map") {
		indexEndOfKey := strings.Index(t, "]")
		keyType, _ := getBasicTypea(t[4:indexEndOfKey])
		valType, _ := getBasicTypea(t[indexEndOfKey+1:])
		return reflect.MapOf(keyType, valType)
	}

	index, err := strconv.Atoi(string(t[1]))
	log.Printf("TTTTTT %d", index)
	ty, _ = getBasicTypea(t[3:])
	return reflect.ArrayOf(index, ty)
}

func getBasicTypea(t string) (ty reflect.Type, e error) {
	switch t {
	case reflect.Int.String(), reflect.Int8.String(), reflect.Int16.String(), reflect.Int32.String(), reflect.Int64.String():
		ty = reflect.TypeOf(int(0))
	case reflect.Uint.String(), reflect.Uint8.String(), reflect.Uint16.String(), reflect.Uint64.String():
		ty = reflect.TypeOf(uint(0))
	case reflect.Float32.String(), reflect.Float64.String():
		ty = reflect.TypeOf(float64(0))
	case reflect.Bool.String():
		ty = reflect.TypeOf(true)
	case reflect.Complex64.String(), reflect.Complex128.String():
		ty = reflect.TypeOf(complex128(0))
	case reflect.String.String():
		ty = reflect.TypeOf("")
	default:
		e = fmt.Errorf("%s is not basic type", t)
	}
	return
}
