package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
)

const (
	cells   = 100
	xyrange = 30
	angle   = math.Pi / 6
)

var width, height int
var xyscale, zscale int
var sin30, cos30 = math.Sin(angle), math.Cos(angle)

//Write is
func Write(writer http.ResponseWriter, w, h int, color string) {
	initQuery(w, h)
	fmt.Fprintf(writer, "<svg xmlns='http://www.w3.org/2000/svg' "+"style='stroke: grey; fill: white; stroke-width: 0.7' "+"width='%d' height='%d'>\n", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			fmt.Fprintf(writer, "<polygon points='%g,%g,%g,%g,%g,%g,%g,%g' stroke='%s'/>\n", ax, ay, bx, by, cx, cy, dx, dy, color)
		}
	}
	fmt.Fprintln(writer, "</svg>")
}

func initQuery(w, h int) {
	width = w
	height = h
	xyscale = w / 2 / xyrange
	zscale = int(float64(height) * float64(0.4))
}

func corner(i, j int) (float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z, b := f(x, y)
	if !b {
		log.Printf("(%f, %f) Invalid Value\n", x, y)
		return 0, 0
	}
	sx := float64(width)/float64(2) + float64(x-y)*cos30*float64(xyscale)
	sy := float64(width)/2 + float64(x+y)*cos30*float64(xyscale) - z*float64(zscale)
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
