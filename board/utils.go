package board

import "math/big"

func PrimalityCheck(number int) bool {
	return big.NewInt(int64(number)).ProbablyPrime(4)
}

var Directions = map[string]Move{
	"right": {1, 0},
	"up":    {0, -1},
	"left":  {-1, 0},
	"down":  {0, 1},
}
