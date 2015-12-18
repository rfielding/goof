package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
//	"math/rand"
	"os"
)

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
	tau        = 2 * math.Pi
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 300
		nframes = 64
		delay   = 8
	)
	//freq := rand.Float64() * tau
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*tau; t += res {
			x := math.Sin(t+phase)*(t*0.01)
			y := math.Cos(t+phase)*(t*0.01)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.2
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
