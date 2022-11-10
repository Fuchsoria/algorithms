package primes

import "testing"

func BenchmarkPrimes(b *testing.B) {
	Cases := Cases[:10]

	for _, bm := range Cases {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = CountPrimes(IsPrime, bm.input)
			}
		})
	}
}

func BenchmarkPrimesSqrt(b *testing.B) {
	Cases := Cases[:10]

	for _, bm := range Cases {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = CountPrimes(IsPrimeSqrt, bm.input)
			}
		})
	}
}
