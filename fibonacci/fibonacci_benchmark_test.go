package fibonacci

import "testing"

func BenchmarkCalculate(b *testing.B) {
	for _, bm := range Cases {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = Calculate(bm.input)
			}
		})
	}
}

func BenchmarkCalculateRecursive(b *testing.B) {
	Cases := Cases[:7]

	for _, bm := range Cases {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = CalculateRecursive(bm.input)
			}
		})
	}
}

func BenchmarkCalculateGold(b *testing.B) {
	// Cases := Cases[:7]

	for _, bm := range Cases {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = CalculateGold(bm.input)
			}
		})
	}
}
