package main

import (
	"fmt"
	"log"
	"math"
)

const (
	width, height = 600, 800
	cells         = 300
	xyrange       = 300
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+"style='stroke: grey; fill: white; stroke-width: 0.7' "+"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			if isPeek(i, j) {
				fmt.Printf("<polygon points='%g,%g,%g,%g,%g,%g,%g,%g' stroke='#ff0000'/>\n", ax, ay, bx, by, cx, cy, dx, dy)
			} else {
				fmt.Printf("<polygon points='%g,%g,%g,%g,%g,%g,%g,%g' stroke='#0000ff'/>\n", ax, ay, bx, by, cx, cy, dx, dy)
			}

		}
	}
	fmt.Println("</svg>")
}

func isPeek(i, j int) bool {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z, _ := f(x, y)
	if z > 0.05 {
		return true
	}
	return false
}

func corner(i, j int) (float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z, b := f(x, y)
	if !b {
		log.Printf("(%f, %f) Invalid Value\n", x, y)
		return 0, 0
	}
	sx := width/2 + (x-y)*cos30*xyscale
	sy := width/2 + (x+y)*cos30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) (float64, bool) {
	r := math.Hypot(x, y) //平方和つまり、距離
	r2 := math.Sin(r) / r
	if math.IsNaN(r2) || math.IsInf(r2, 0) {
		return 0, false
	}
	return r2, true
}
