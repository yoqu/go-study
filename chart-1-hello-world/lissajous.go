package main

import (
	"io"
	"math/rand"
	"image/gif"
	"image"
	"image/color"
	"math"
	"os"
	"fmt"
)

var palette = []color.Color{color.White, color.Black, color.RGBA{102, 255, 102, 255}}

const (
	whiteIndex = 0
	blackIndex = 1
	greenIndex = 2
)

func main() {
	file,error :=os.Create("test.gif")
	if error!=nil{
		fmt.Printf(" open file error")
		return;
	}
	defer file.Close()
	//lissajous(os.Stdout)
	lissajous(file)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
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
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), greenIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)

}