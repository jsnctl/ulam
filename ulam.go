package main

import (
	"image/color"
	"math/big"
)

func PrimalityCheck(number int) bool {
	return big.NewInt(int64(number)).ProbablyPrime(4)
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

func (board *Board) UlamSpiral() {
	cursor := Cursor{
		number:  0,
		x:       board.centre.X,
		y:       board.centre.Y,
		isPrime: false,
	}

	moves := GenerateMoves(board.width * board.height)
	for i, move := range moves {
		cursor.moveCursor(i+1, move)
		if cursor.isPrime {
			board.img.Set(cursor.x, cursor.y, color.Black)
		}
	}
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
