package main

import (
	"fmt"
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

	cursor := Cursor{
		number:  0,
		x:       centre.X,
		y:       centre.Y,
		isPrime: false,
	}

	for i := 0; i < 200; i++ {
		cursor.moveCursor(i)
		img.Set(cursor.x, cursor.y, color.Black)
	}

	f, _ := os.Create("output.png")
	png.Encode(f, img)
}

type Cursor struct {
	number  int
	x       int
	y       int
	isPrime bool
}

type Move struct {
	xDelta int
	yDelta int
}

var directions = map[string]Move{
	"right": {1, 0},
	"up":    {0, -1},
	"left":  {-1, 0},
	"down":  {0, 1},
}

func GenerateSeries(steps int) []int {
	series := make([]int, 0)
	total := 0
	for step := 1; step < steps+1; step++ {
		diff := (step + 1) / 2
		total += diff

		if total >= steps {
			series = append(series, step-diff)
			break
		}

		series = append(series, diff)
	}
	return series
}

func GenerateMoves(steps int) []string {
	result := make([]string, 0)
	for step := 0; step < steps; step++ {
		result = append(result, "right")
	}
	fmt.Println(result)
	return []string{"right"}
}

func (c *Cursor) moveCursor(number int) {
	c.number = number
	c.x += 1
	c.y += 1
}
