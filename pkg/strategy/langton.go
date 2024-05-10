package strategy

import (
	"github.com/jsnctl/ulam/pkg/board"
	"image/color"
	"log/slog"
)

func turnRight(direction string) string {
	lookup := map[string]string{
		"right": "down",
		"down":  "left",
		"left":  "up",
		"up":    "right",
	}
	return lookup[direction]
}
func turnLeft(direction string) string {
	lookup := map[string]string{
		"right": "up",
		"down":  "right",
		"left":  "down",
		"up":    "left",
	}
	return lookup[direction]
}

func compareColors(left color.Color, right color.Color) bool {
	leftA, leftB, leftC, leftD := left.RGBA()
	rightA, rightB, rightC, rightD := right.RGBA()
	if (leftA == rightA) && (leftB == rightB) && (leftC == rightC) && (leftD == rightD) {
		return true
	}
	return false
}

func (s *Strategy) LangtonAnt(steps int) {
	slog.Info("Running Langton ant...")
	cursor := board.Cursor{
		Number: 0,
		X:      s.Board.Centre.X,
		Y:      s.Board.Centre.Y,
		Facing: "up",
	}

	for i := 0; i < steps; i++ {
		cursor.Colour = s.Board.GetColour(cursor.X, cursor.Y)
		if compareColors(cursor.Colour, color.White) {
			cursor.Facing = turnRight(cursor.Facing)
			s.Board.Img.Set(cursor.X, cursor.Y, color.Black)
			cursor.MoveCursor(i+1, board.Directions[cursor.Facing])
		} else {
			cursor.Facing = turnLeft(cursor.Facing)
			s.Board.Img.Set(cursor.X, cursor.Y, color.White)
			cursor.MoveCursor(i+1, board.Directions[cursor.Facing])
		}
	}
}
