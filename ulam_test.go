package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerateSeries(t *testing.T) {
	//assert.Equal(t, []int{1, 1, 2, 2, 2}, GenerateSeries(8))
	assert.Equal(t, []int{1, 1, 2, 2, 3, 3, 1}, GenerateSeries(13))
}

func TestGenerateMoves(t *testing.T) {
	assert.Equal(t, []string{"right"}, GenerateMoves(10))
	//assert.Equal(t, []string{"right", "up"}, GenerateMoves(2))
}
