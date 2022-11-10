package primes

import (
	"math"
)

var Cases = ParseCases()

type PrimeFunc = func(n int) bool

func CountPrimes(f PrimeFunc, n int) int {
	count := 0

	for i := 2; i <= n; i++ {
		if f(i) {
			count++
		}
	}

	return count
}

func IsPrime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func IsPrimeSqrt(n int) bool {
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}

	for i := 3; float64(i) <= math.Sqrt(float64(n)); i += 2 {
		if n%i == 0 {
			return false
		}
	}

	return true
}
