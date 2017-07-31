// Copyright © 2017 Ryutarou Ono.

package main

import (
	"errors"
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"os"
)

func main() {
	t := flag.String("t", "jpeg", "to set output type")
	flag.Parse()

	in := os.Stdin
	out := os.Stdout
	var err error
	switch *t {
	case "jpeg", "JPEG":
		err = toJPEG(in, out)
	case "gif", "GIF":
		err = toGIF(in, out)
	case "png", "PNG":
		err = toPNG(in, out)
	default:
		err = errors.New("Unsupported Type")
	}
	if err != nil {
		log.Println(err)
	}
}

func toJPEG(in io.Reader, out io.Writer) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)
	return jpeg.Encode(out, img, &jpeg.Options{95})
}

func toPNG(in io.Reader, out io.Writer) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)
	return png.Encode(out, img)
}

func toGIF(in io.Reader, out io.Writer) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintln(os.Stderr, "Input format =", kind)
	return gif.Encode(out, img, &gif.Options{45, nil, nil})
}
