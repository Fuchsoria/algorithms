package exponentiation

import "math"

var Cases = ParseCases()

func Calculate(a float64, n float64) float64 {
	v := 1.0

	for i := 0; i < int(n); i++ {
		v *= a
	}

	return v
}

func CalculateMath(a float64, n float64) float64 {
	return math.Pow(a, n)
}
