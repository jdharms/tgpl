// Exercise 3.1: If the function f returns a non-finite float64 value,
// the SVG file will contain invalid <polygon> elements (although many
// SVG renderers handle this gracefully). Modify the program to skip invalid polygons.

// Exercise 3.3: Color each polygon based on its height, so that the
// peaks are colored red (#ff0000) and the valleys blue (#0000ff).

// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"fmt"
	"math"
	"strconv"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, az, aok := corner(i+1, j)
			bx, by, bz, bok := corner(i, j)
			cx, cy, cz, cok := corner(i, j+1)
			dx, dy, dz, dok := corner(i+1, j+1)
			if aok && bok && cok && dok {
				polyHeight := max(az, bz, cz, dz)
				scaledHeight := polyHeight * 255 * 4

				var colorStr string
				if scaledHeight > 0 {
					color := scaledHeight
					colorStr = "rgb(" + strconv.Itoa(int(color)) + ", 0, 0)"
				} else {
					color := -scaledHeight
					colorStr = "rgb(0, 0, " + strconv.Itoa(int(color)) + ")"
				}
				fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' style='stroke:"+colorStr+"'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy)
			}
		}
	}
	fmt.Println("</svg>")
}

func max(a float64, b float64, c float64, d float64) float64 {
	result := math.Max(a, b)
	result = math.Max(result, c)
	result = math.Max(result, d)
	return result
}

func corner(i, j int) (float64, float64, float64, bool) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z, ok := f(x, y)

	if ok {
		// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
		sx := width/2 + (x-y)*cos30*xyscale
		sy := height/2 + (x+y)*sin30*xyscale - z*zscale
		return sx, sy, z, true
	}
	return 0, 0, 0, false
}

func f(x, y float64) (float64, bool) {
	r := math.Hypot(x, y) // distance from (0,0)
	result := math.Sin(r) / r
	if math.IsNaN(result) {
		return 0, false
	}

	return result, true
}
