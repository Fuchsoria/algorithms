package strs

import (
	"testing"
)

func BenchmarkBM(b *testing.B) {
	full := "ABAAABCDLHEREABCD"
	part := "ABC"

	b.Run("search BM default", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			bm := NewBM()

			bm.search(full, part)
		}
	})

	b.Run("search BM 1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			bm := NewBMUpdated()

			bm.Find(full, part)
		}
	})

	b.Run("search BM 2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			bm := NewBMUpdated()

			bm.Find2(full, part)
		}
	})

	b.Run("search BM 3", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			bm := NewBMUpdated()

			bm.Find3(full, part)
		}
	})
}
