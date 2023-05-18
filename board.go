package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

type Board struct {
	img         *image.RGBA
	width       int
	height      int
	topLeft     image.Point
	bottomRight image.Point
	centre      image.Point
}

func NewBoard(width int, height int) *Board {
	board := Board{width: width, height: height}
	board.topLeft = image.Point{X: 0, Y: 0}
	board.bottomRight = image.Point{X: width, Y: height}
	board.centre = image.Point{X: width / 2, Y: height / 2}
	board.img = image.NewRGBA(image.Rectangle{Min: board.topLeft, Max: board.bottomRight})
	return &board
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

func main() {
	board := NewBoard(1000, 1000)

	draw.Draw(
		board.img,
		board.img.Bounds(),
		&image.Uniform{C: color.RGBA{R: 255, G: 255, B: 255, A: 255}},
		image.ZP,
		draw.Src,
	)

	board.UlamSpiral()

	f, _ := os.Create("output.png")
	png.Encode(f, board.img)
}
