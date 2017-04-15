package anime

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
)

var palette = []color.Color{color.RGBA{0x00, 0x00, 0x00, 0x00}, color.RGBA{0x00, 0xff, 0x00, 0x00}}

const (
	whiteIndex = 0
	blackIndex = 1
)

// Lissajous is write lissajous
// pallet {background color, line color}
func Lissajous(out io.Writer, palette []color.Color) {
	const (
		cycles  = 2
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 3
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		side := 2*size + 1
		rect := image.Rect(0, 0, side, side)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
