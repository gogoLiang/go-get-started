package main

import (
	"image"
	"image/color"
	"image/gif"
)

func main() {
	gif := gif.GIF{LoopCount: 64}

	for i := 0; i < 64; i++ {
		rect := image.Rect(0, 0, 101, 101)
		image.NewPaletted(rect, []color.Color{color.Black, color.White})
	}
}
