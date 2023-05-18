package main

import (
	"image/color"
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

func (board *Board) LangtonAnt(steps int) {
	cursor := Cursor{
		number: 0,
		x:      board.centre.X,
		y:      board.centre.Y,
		facing: "up",
	}

	for i := 0; i < steps; i++ {
		cursor.colour = board.getColour(cursor.x, cursor.y)
		if compareColors(cursor.colour, color.White) {
			cursor.facing = turnRight(cursor.facing)
			board.img.Set(cursor.x, cursor.y, color.Black)
			cursor.moveCursor(i+1, directions[cursor.facing])
		} else {
			cursor.facing = turnLeft(cursor.facing)
			board.img.Set(cursor.x, cursor.y, color.White)
			cursor.moveCursor(i+1, directions[cursor.facing])
		}
	}
}
