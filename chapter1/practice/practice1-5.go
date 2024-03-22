package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"time"
)

var palette = []color.Color{
	color.Black,
	color.RGBA{R: 73, G: 156, B: 84, A: 1},
	color.RGBA{R: 1, G: 158, B: 208, A: 1},
	color.RGBA{R: 220, G: 50, B: 90, A: 1},
	color.RGBA{R: 73, G: 16, B: 184, A: 1},
	color.RGBA{R: 23, G: 116, B: 124, A: 1},
	color.White,
}

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	handler := func(w http.ResponseWriter, r *http.Request) {
		lissajous(w)
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
	return
	//lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 500
		res     = 5
		size    = 100
		nframes = 100
		delay   = 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)

			randIndex := rand.Intn(len(palette))
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(randIndex))

		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
