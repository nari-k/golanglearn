package main

import (
	"fmt"
	"io"
	"math"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	angle         = math.Pi / 6
)

type params struct {
	width   float64
	height  float64
	color   string
	xyscale float64
	zscale  float64
}

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func surface(out io.Writer, p *params) {
	if p.width < 1 {
		p.width = width
	}
	if p.height < 1 {
		p.height = height
	}
	p.xyscale = p.width / 2 / xyrange
	p.zscale = p.height * 0.4
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: %s; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", p.color, p.width, p.height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, az := corner(i+1, j, p)
			bx, by, bz := corner(i, j, p)
			cx, cy, cz := corner(i, j+1, p)
			dx, dy, dz := corner(i+1, j+1, p)
			if isInf([]float64{ax, ay, bx, by, cx, cy, dx, dy}) {
				fmt.Fprintf(out, "<polygon points='%g,%g,%g,%g,%g,%g,%g,%g' fill='%s'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy, color(az, bz, cz, dz))
			}
		}
	}
	fmt.Fprintf(out, "</svg>")
}

func color(az, bz, cz, dz float64) string {
	z := ((az+bz+cz+dz)/4 - 0.3) * 1.7
	color := int(0xff * (math.Abs(z)))
	if color > 0xff {
		color = 0xff
	}
	green := 0xff - color
	base := "#"
	rgb := make([]byte, 0, 7)
	rgb = append(rgb, base...)

	strColor := fmt.Sprintf("%02x", color)
	grnColor := fmt.Sprintf("%02x", green)
	blk := "00"
	if z > 0 {
		rgb = append(rgb, strColor...)
		rgb = append(rgb, grnColor...)
		rgb = append(rgb, blk...)
	} else {
		rgb = append(rgb, blk...)
		rgb = append(rgb, grnColor...)
		rgb = append(rgb, strColor...)
	}
	return string(rgb)
}

func isInf(fa []float64) bool {
	for _, f := range fa {
		if math.IsInf(f, 0) {
			return false
		}
	}
	return true
}

func corner(i, j int, p *params) (float64, float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z := f(x, y)
	sx := p.width/2 + (x-y)*cos30*p.xyscale
	sy := p.height/2 + (x+y)*sin30*p.xyscale - z*p.zscale
	return sx, sy, z
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}
