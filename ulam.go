package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	width := 100
	height := 100
	topLeft := image.Point{X: 0, Y: 0}
	bottomRight := image.Point{X: width, Y: height}
	centre := image.Point{X: width / 2, Y: height / 2}

	img := image.NewRGBA(image.Rectangle{Min: topLeft, Max: bottomRight})
	img.Set(centre.X, centre.Y, color.Black)

	f, _ := os.Create("output.png")
	png.Encode(f, img)
}
