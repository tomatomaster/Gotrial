#!/usr/bin/env bash
go build jpeg.go
go build mandelbrot.go
./mandelbrot| ./jpeg -t=gif > mandelbrot.gif
./mandelbrot| ./jpeg -t=png > mandelbrot.png
./mandelbrot| ./jpeg -t=jpeg > mandelbrot.jpeg