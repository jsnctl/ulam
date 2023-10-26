package board

import (
	"image"
	"image/color"
)

type Board struct {
	Img         *image.RGBA
	Width       int
	Height      int
	TopLeft     image.Point
	BottomRight image.Point
	Centre      image.Point
}

func NewBoard(width int, height int) *Board {
	board := Board{Width: width, Height: height}
	board.TopLeft = image.Point{X: 0, Y: 0}
	board.BottomRight = image.Point{X: width, Y: height}
	board.Centre = image.Point{X: width / 2, Y: height / 2}
	board.Img = image.NewRGBA(image.Rectangle{Min: board.TopLeft, Max: board.BottomRight})
	return &board
}

type Cursor struct {
	Number  int
	X       int
	Y       int
	IsPrime bool
	Colour  color.Color
	Facing  string
}

func (board *Board) GetColour(x int, y int) color.Color {
	return board.Img.At(x, y)
}

func (c *Cursor) MoveCursor(number int, move Move) {
	c.Number = number
	c.IsPrime = PrimalityCheck(c.Number)
	c.X += move.XDelta
	c.Y += move.YDelta
}

type Move struct {
	XDelta int
	YDelta int
}
