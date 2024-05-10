package main

import (
	"github.com/jsnctl/ulam/pkg/board"
	"github.com/jsnctl/ulam/pkg/strategy"
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

	var which string
	if len(os.Args) > 1 {
		which = os.Args[1]
	}

	switch which {
	case "ulam":
		s.UlamSpiral()
	case "langton":
		s.LangtonAnt(1e6)
	case "":
		s.UlamSpiral()
		s.LangtonAnt(1e5)
	}

	f, _ := os.Create("output.png")
	png.Encode(f, s.Board.Img)
}
