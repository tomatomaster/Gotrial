package main

import (
	"fmt"
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
		width, height          = 1024, 1024
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

func root(images ...color.Color) color.RGBA {
	var r, g, b, a uint32
	for _, image := range images {
		_r, _g, _b, _a := image.RGBA()
		fmt.Print(_r, _g, _b, _a)
		r = (r + _r) / 2
		g = (g + _g) / 2
		b = (b + _b) / 2
		a = (a + _a) / 2
	}
	return color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a)}
}

func mandelbrot(z complex128) color.Color {
	const iterations = 51
	const criterior = 0.00000000000000000001
	const contrast = 5

	//iterationsの回数が明確なので、0-256の範囲で十分
	//200回以内に発散しなければ収束しているとみなしているが、根拠が良くわからない
	for n := uint8(0); n < iterations; n++ {
		//v = v*v + z //zk+1 = zkn + C　（n=2, C=z)

		z = newton(z)
		if cmplx.Abs((1+0i)-z) < criterior {
			return color.RGBA{255 - contrast*n, 0, 0, 255 - contrast*n}
		} else if cmplx.Abs((-1+0i)-z) < criterior {
			return color.RGBA{0, 255 - contrast*n, contrast * n, 255 - contrast*n}
		} else if cmplx.Abs((0+1i)-z) < criterior {
			return color.RGBA{0, 0, 255 - contrast*n, 255}
		} else if cmplx.Abs((0-1i)-z) < criterior {
			return color.RGBA{100 - contrast*n, 100 - contrast*n, 100 - contrast*n, 255}
		}
	}
	return color.RGBA{200, 200, 50, 255} //収束
}

func newton(x complex128) complex128 {
	return x - (x*x*x*x-1.0)/(4*x*x*x)
}
