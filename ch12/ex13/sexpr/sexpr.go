// Copyright © 2017 Ryutarou Ono.

package sexpr

import (
	"bytes"
	"fmt"
	"reflect"
)

const tab = "        "

func Marshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	if err := encode(&buf, reflect.ValueOf(v)); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func encode(buf *bytes.Buffer, v reflect.Value) error {
	switch v.Kind() {
	case reflect.Invalid:
		buf.WriteString("nil")

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fmt.Fprintf(buf, "%d", v.Int())

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		fmt.Fprintf(buf, "%d", v.Uint())

	case reflect.Float32, reflect.Float64:
		fmt.Fprintf(buf, "%f", v.Float())

	case reflect.String:
		fmt.Fprintf(buf, "%q", v.String())

	case reflect.Bool:
		if v.Bool() {
			fmt.Fprintf(buf, "%s", "t")
		}
		fmt.Fprintf(buf, "%s", "nil")

	case reflect.Ptr:
		return encode(buf, v.Elem())

	case reflect.Array, reflect.Slice:
		buf.WriteByte('(')
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				buf.WriteByte(' ')
			}
			if i != 0 {
				fmt.Fprint(buf, tab)
			}
			if err := encode(buf, v.Index(i)); err != nil {
				return err
			}
			if i != v.Len()-1 {
				buf.WriteByte('\n')
			}
		}
		buf.WriteByte(')')

	case reflect.Struct:
		buf.WriteByte('(')
		for i := 0; i < v.NumField(); i++ {
			if i > 0 {
				buf.WriteByte(' ')
			}
			if tagname := getTagName(v.Type().Field(i)); tagname != "" { //Add ex13
				fmt.Fprintf(buf, "(%s ", tagname)
			} else {
				fmt.Fprintf(buf, "(%s ", v.Type().Field(i).Name)
			}
			if err := encode(buf, v.Field(i)); err != nil {
				return err
			}
			buf.WriteByte(')')
			buf.WriteByte('\n')
		}
		buf.WriteByte(')')

	case reflect.Map:
		buf.WriteByte('(')
		for i, key := range v.MapKeys() {
			if i > 0 {
				buf.WriteByte(' ')
			}
			if i != 0 {
				fmt.Fprint(buf, tab)
			}
			buf.WriteByte('(')
			if err := encode(buf, key); err != nil { //Key (key val)
				return err
			}
			buf.WriteByte(' ')
			if err := encode(buf, v.MapIndex(key)); err != nil { //Val(key val)
				return err
			}
			buf.WriteByte(')')
			if i != v.Len()-1 {
				buf.WriteByte('\n')
			}
		}
		buf.WriteByte(')')

	case reflect.Complex64, reflect.Complex128:
		fmt.Fprintf(buf, "#C(%f %f)", real(v.Complex()), imag(v.Complex()))

	case reflect.Interface:
		fmt.Fprintf(buf, "%#v", v.Interface())

	default: //chan, func
		return fmt.Errorf("unsupported type: %s", v.Type())
	}
	return nil
}

func getTagName(v reflect.StructField) string {
	tag := v.Tag
	name := tag.Get("sexpr")
	return name
}
