package fibonacci

import (
	"math"
)

var Cases = ParseCases()

func Calculate(n float64) float64 {
	if n <= 1 {
		return n
	}

	var a, b float64 = 0, 1

	for i := 2; i < int(n); i++ {
		f := a + b
		a = b
		b = f
	}

	return a + b
}

func CalculateRecursive(n float64) float64 {
	if n <= 1 {
		return n
	}

	return CalculateRecursive(n-1) + CalculateRecursive(n-2)
}

func CalculateGold(n float64) float64 {
	if n <= 1 {
		return n
	}

	sqrt5 := math.Sqrt(5)
	f := (sqrt5 + 1) / 2

	return math.Round(math.Pow(f, n) / sqrt5)
}
