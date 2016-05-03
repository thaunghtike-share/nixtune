package main

import (
	"fmt"
	"math"
	"os"
)

const (
	goldenRatio = 1.6180
	width       = 500 * goldenRatio
	height      = 500                 // canvas size in pixels
	cells       = 100                 // number of grid cells
	xyrange     = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale     = width / 2 / xyrange // pixels per x or y unit
	zscale      = height * 0.4        // pixels per z unit
	angle       = math.Pi / 6         // angle of x, y axes (=30°)
	logoName    = "acksin autotune"
)

var (
	colors = []string{
		"tomato",
		"gold",
		"lightsteelblue",
	}
	sin30 = math.Sin(angle)
	cos30 = math.Cos(angle)
)

func main() {
	version := os.Args[1]

	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' style='stroke: orangered; fill: black; stroke-width: 0.2' width='%dpx' height='%dpx'>\n", int64(width), height)

	colorIndex := 0

	for i := 0; i < cells; i++ {
		colorIndex++
		if colorIndex == len(colors) {
			colorIndex = 0
		}
		color := colors[colorIndex]

		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)

			fmt.Printf("<polygon fill='%s' points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				color, ax, ay, bx, by, cx, cy, dx, dy)
		}
	}

	fmt.Printf(`<text text-anchor="middle" x="%f" y="%f" font-size="64" fill="darkslategray" font-family="helvetica"><tspan dy="1.2em">%s</tspan><tspan dy="1.2em">%s</tspan></text>`, float64(width/2), float64(height/2), logoName, version)
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := math.Sin(y*x/50) * math.Cos(x*y/50) / 2.71

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}
