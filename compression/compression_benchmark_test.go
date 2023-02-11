package compression

import (
	"testing"
)

func BenchmarkHashtable(b *testing.B) {
	input := "wwwwaaadexxxxxxywww"
	output := "4w3a1d1e6x1y3w"

	b.Run("encode", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			c := NewCompression()

			for i := 0; i < 10000; i++ {
				c.Encode(input)
			}
		}
	})

	b.Run("decode", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			c := NewCompression()

			for i := 0; i < 10000; i++ {
				c.Decode(output)
			}
		}
	})
}
