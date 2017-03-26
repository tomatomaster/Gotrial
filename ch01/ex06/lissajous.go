package main

import (
	"image/color"
	"math/rand"
	"os"
	"time"

	"./anime"
)

var palette = []color.Color{
	color.RGBA{0x00, 0x00, 0x00, 0xff},
	color.RGBA{0x00, 0xff, 0x00, 0xff},
	color.RGBA{0xff, 0x00, 0x00, 0xff}}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	anime.Lissajous(os.Stdout, palette)
}
