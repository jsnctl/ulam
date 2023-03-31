package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math/big"
	"os"
)

func main() {
	width := 1000
	height := 1000
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

	moves := GenerateMoves(width * height)
	for i, move := range moves {
		cursor.moveCursor(i+1, move)
		if cursor.isPrime {
			img.Set(cursor.x, cursor.y, color.Black)
		}
	}

	f, _ := os.Create("output.png")
	png.Encode(f, img)
}

func PrimalityCheck(number int) bool {
	return big.NewInt(int64(number)).ProbablyPrime(4)
}

type Cursor struct {
	number  int
	x       int
	y       int
	isPrime bool
}

func (c *Cursor) moveCursor(number int, move Move) {
	c.number = number
	c.isPrime = PrimalityCheck(c.number)
	c.x += move.xDelta
	c.y += move.yDelta
}

type Move struct {
	xDelta int
	yDelta int
}

type DirectionLoop struct {
	Directions []string
	cursor int
}

func NewDirectionLoop() *DirectionLoop {
	loop := &DirectionLoop{}
	loop.Directions = []string{"right", "up", "left", "down"}
	return loop
}

func (loop *DirectionLoop) GetNextDirection() string {
	value := loop.Directions[loop.cursor]
	loop.cursor += 1
	if loop.cursor > 3 {
		loop.cursor = 0
	}
	return value
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

		if total+diff >= steps {
			series = append(series, steps-total)
			break
		}

		total += diff
		series = append(series, diff)
	}
	return series
}

func GenerateMoves(steps int) []Move {
	series := GenerateSeries(steps)
	result := make([]Move, 0)
	loop := NewDirectionLoop()
	for _, step := range series {
		chunk := GenerateMoveChunk(step, loop.GetNextDirection())
		result = append(result, chunk...)
	}
	return result
}

func GenerateMoveChunk(size int, value string) []Move {
	chunk := make([]Move, size)
	for i := 0; i < size; i++ {
		chunk[i] = directions[value]
	}
	return chunk
}

