/*
* Модифицируйте функцию corner из gopl.io/ch3/surface
* (раздел 3), чтобы она использовала именованные результаты
* и инструкцию пустого возврата.
 */

package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
)

const (
	width, height = 1200, 800
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "web" {
		handler := func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "image/svg+xml")
			render(w, width, height, cells)
		}
		http.HandleFunc("/", handler)

		log.Fatal(http.ListenAndServe("localhost:8000", nil))
	}
}

func corner(i, j int) (sx, sy float64, ok bool) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// z := f(x, y)
	// z := eggBox(x, y)
	z := saddle(x, y)

	sx = width/2 + (x-y)*cos30*xyscale
	sy = height/2 + (x+y)*sin30*xyscale - z*zscale
	ok = math.IsInf(sx, 0) || math.IsInf(sy, 0) || math.IsNaN(sx) || math.IsNaN(sy)

	return
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r)
}

func eggBox(x, y float64) float64 {
	return 0.2 * (math.Cos(x) + math.Cos(y))
}

func saddle(x, y float64) float64 {
	a := 25.0
	b := 17.0
	a2 := a * a
	b2 := b * b
	return (y*y/a2 - x*x/b2)
}

func render(w io.Writer, width, height, cells int) {
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='"+strconv.Itoa(width)+"' "+
		"height='"+strconv.Itoa(height)+"'>\n")
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, incorrect := corner(i+1, j)
			if incorrect {
				continue
			}
			bx, by, incorrect := corner(i, j)
			if incorrect {
				continue
			}
			cx, cy, incorrect := corner(i, j+1)
			if incorrect {
				continue
			}
			dx, dy, incorrect := corner(i+1, j+1)
			if incorrect {
				continue
			}
			fmt.Fprintf(w, "<polygon points='"+
				strconv.FormatFloat(ax, 'g', 14, 64)+","+strconv.FormatFloat(ay, 'g', 14, 64)+" "+
				strconv.FormatFloat(bx, 'g', 14, 64)+","+strconv.FormatFloat(by, 'g', 14, 64)+" "+
				strconv.FormatFloat(cx, 'g', 14, 64)+","+strconv.FormatFloat(cy, 'g', 14, 64)+" "+
				strconv.FormatFloat(dx, 'g', 14, 64)+","+strconv.FormatFloat(dy, 'g', 14, 64)+"'/>\n")
		}
	}
	fmt.Fprintf(w, "</svg>")
}
