package strs

import (
	"testing"
)

func BenchmarkKMP(b *testing.B) {
	full := "ABAAABCDLHEREABCD"
	part := "ABC"

	b.Run("search KMP default", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			kmp := NewKMP()

			kmp.Search(full, part)
		}
	})
}
