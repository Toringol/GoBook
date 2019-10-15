/*
* Измените сервер с фигурами Лиссажу так, чтобы значения параметров считывались из URL.
* Например, URL вида http://localhost:8000/?cycles=20 устанавливает количество циклов
* равным 20 вместо значения по умолчанию, равного 5. Используйте функцию strconv.Atoi для
* преобразования строкового параметра в целое число. Просмотреть документацию по данной
* функции можно с помощью go doc strconv.Atoi.
 */

// Lissajous generates GIF animations of random Lissajous figures.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

import (
	"log"
	"net/http"
	"strconv"
	"time"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

type DataFigure struct {
	cycles  float64 // number of complete x oscillator revolutions
	res     float64 // angular resolution
	size    int     // image canvas covers [-size..+size]
	nframes int     // number of animation frames
	delay   int     // delay between frames in 10ms units
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	dataFigure := DataFigure{
		cycles:  5,
		res:     0.001,
		size:    100,
		nframes: 64,
		delay:   8,
	}

	if len(os.Args) > 1 && os.Args[1] == "web" {
		handler := func(w http.ResponseWriter, r *http.Request) {
			if err := r.ParseForm(); err != nil {
				log.Print(err)
			}
			for key, value := range r.Form {
				changeDefaultDataFigure(key, value[0], &dataFigure)
			}
			lissajous(w, dataFigure)
		}
		http.HandleFunc("/", handler)

		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	lissajous(os.Stdout, dataFigure)
}

func lissajous(out io.Writer, dataFigure DataFigure) {
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: dataFigure.nframes}
	phase := 0.0 // phase difference
	for i := 0; i < dataFigure.nframes; i++ {
		rect := image.Rect(0, 0, 2*dataFigure.size+1, 2*dataFigure.size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < dataFigure.cycles*2*math.Pi; t += dataFigure.res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(dataFigure.size+int(x*float64(dataFigure.size)+0.5), dataFigure.size+int(y*float64(dataFigure.size)+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, dataFigure.delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

func changeDefaultDataFigure(key string, value string, dataFigure *DataFigure) {
	var err error
	if key == "cycles" {
		dataFigure.cycles, err = strconv.ParseFloat(value, 64)
	}
	if key == "res" {
		dataFigure.res, err = strconv.ParseFloat(value, 64)
	}
	if key == "size" {
		dataFigure.size, err = strconv.Atoi(value)
	}
	if key == "nframes" {
		dataFigure.nframes, err = strconv.Atoi(value)
	}
	if key == "delay" {
		dataFigure.delay, err = strconv.Atoi(value)
	}
	if err != nil {
		log.Println(err)
	}
}
