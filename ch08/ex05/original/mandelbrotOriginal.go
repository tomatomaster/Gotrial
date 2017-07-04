package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		//http://azisava.sakura.ne.jp/mandelbrot/definition.html
		//2以上の範囲では必ず発散する
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 4096, 4096
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y) //複素数型
			// image
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img)
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	//iterationsの回数が明確なので、0-256の範囲で十分
	//200回以内に発散しなければ収束しているとみなしているが、根拠が良くわからない
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z           //zk+1 = zkn + C　（n=2, C=z)
		if cmplx.Abs(v) > 2 { //C>2は必ず発散する
			if n < 10 {
				return color.RGBA{n * 25, 255 - n*25, 0, 255}
			} else if 20 < n && n < 100 {
				return color.RGBA{n * 10, n * 10, 255 - (n-20)*20, 255}
			} else if 100 < n {
				return color.RGBA{255 - (n-100)*2, n * 50, 0, 255}
			}
		}
	}
	return color.RGBA{200, 200, 50, 255}
}
