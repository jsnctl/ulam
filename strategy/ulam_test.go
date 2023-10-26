package strategy

import (
	"github.com/jsnctl/ulam/board"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateSeries(t *testing.T) {
	assert.Equal(t, []int{1, 1, 2, 2, 2}, GenerateSeries(8))
	assert.Equal(t, []int{1, 1, 2, 2, 3, 3, 1}, GenerateSeries(13))
	assert.Equal(t, []int{1, 1, 2, 2, 3, 3, 4, 4, 1}, GenerateSeries(21))
}

func TestGenerateMoveChunk(t *testing.T) {
	assert.Equal(t, []board.Move{{0, -1}}, GenerateMoveChunk(1, "up"))
	assert.Equal(t, []board.Move{{-1, 0}, {-1, 0}}, GenerateMoveChunk(2, "left"))
	assert.Equal(t, []board.Move{{1, 0}, {1, 0}, {1, 0}, {1, 0}}, GenerateMoveChunk(4, "right"))
}

func TestDirectionLoop(t *testing.T) {
	loop := NewDirectionLoop()
	assert.Equal(t, "right", loop.GetNextDirection())
	assert.Equal(t, "up", loop.GetNextDirection())
	assert.Equal(t, "left", loop.GetNextDirection())
	assert.Equal(t, "down", loop.GetNextDirection())
	assert.Equal(t, "right", loop.GetNextDirection())
}

func TestGenerateMoves(t *testing.T) {
	moves := GenerateMoves(3)
	assert.Equal(t,
		[]board.Move{{1, 0}, {0, -1}, {-1, 0}},
		moves,
	)
}
