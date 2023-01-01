package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

func main() {
	width := 500
	height := 500
	topLeft := image.Point{X: 0, Y: 0}
	bottomRight := image.Point{X: width, Y: height}
	centre := image.Point{X: width / 2, Y: height / 2}

	img := image.NewRGBA(image.Rectangle{Min: topLeft, Max: bottomRight})
	draw.Draw(img, img.Bounds(), &image.Uniform{C: color.RGBA{R: 255, G: 255, B: 255, A: 255}}, image.ZP, draw.Src)
	img.Set(centre.X, centre.Y, color.Black)

	f, _ := os.Create("output.png")
	png.Encode(f, img)
}
