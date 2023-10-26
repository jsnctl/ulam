package strategy

import (
	"github.com/jsnctl/ulam/board"
	"image/color"
)

func (s *Strategy) UlamSpiral() {
	cursor := board.Cursor{
		Number:  0,
		X:       s.Board.Centre.X,
		Y:       s.Board.Centre.Y,
		IsPrime: false,
	}

	moves := GenerateMoves(s.Board.Width * s.Board.Height)
	for i, move := range moves {
		cursor.MoveCursor(i+1, move)
		if cursor.IsPrime {
			s.Board.Img.Set(cursor.X, cursor.Y, color.Black)
		}
	}
}

func GenerateMoves(steps int) []board.Move {
	series := GenerateSeries(steps)
	result := make([]board.Move, 0)
	loop := NewDirectionLoop()
	for _, step := range series {
		chunk := GenerateMoveChunk(step, loop.GetNextDirection())
		result = append(result, chunk...)
	}
	return result
}

func GenerateMoveChunk(size int, value string) []board.Move {
	chunk := make([]board.Move, size)
	for i := 0; i < size; i++ {
		chunk[i] = board.Directions[value]
	}
	return chunk
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

type DirectionLoop struct {
	Directions []string
	cursor     int
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
