package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			if isInf([]float64{ax, ay, bx, by, cx, cy, dx, dy}) {
				fmt.Printf("<polygon points='%g,%g,%g,%g,%g,%g,%g,%g'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy)
			}
		}
	}
	fmt.Println("</svg>")
}

func isInf(fa []float64) bool {
	for _, f := range fa {
		if math.IsInf(f, 0) {
			return false
		}
	}
	return true
}

func corner(i, j int) (float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z := eggbox(x, y)
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func eggbox(x, y float64) float64 {
	f, g := 2.0, 20.0
	return math.Pow(f, math.Sin(y)) * math.Pow(f, math.Sin(x)) / g
}

func saddle(x, y float64) float64 {
	h := 2.3
	return -math.Cos(x/xyrange*h) + math.Cos(y/xyrange*h)
}
