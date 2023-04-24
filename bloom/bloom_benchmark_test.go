package bloom

import (
	"testing"
)

func BenchmarkBloomFilter(b *testing.B) {
	b.Run("Filter", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			bloom := NewBloomFilter(100)

			for _, word := range words {
				bloom.Insert(word)
			}
		}
	})
}
