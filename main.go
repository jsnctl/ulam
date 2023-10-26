package main

import (
	"github.com/jsnctl/ulam/board"
	"github.com/jsnctl/ulam/strategy"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)

func main() {
	b := board.NewBoard(1000, 1000)

	draw.Draw(
		b.Img,
		b.Img.Bounds(),
		&image.Uniform{C: color.White},
		image.ZP,
		draw.Src,
	)

	s := strategy.Strategy{Board: b}
	s.LangtonAnt(1000)
	s.UlamSpiral()

	f, _ := os.Create("output.png")
	png.Encode(f, s.Board.Img)
}
