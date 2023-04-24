package barn

import (
	"testing"
)

func BenchmarkBarn(b *testing.B) {
	b.Run("Solve N12", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			N := 12
			F := 100
			Seed := 123456

			barn := NewBarn(N)
			barn.generateMap(int64(F), int64(Seed))

			barn.Solve()
		}
	})

	b.Run("Solve N100", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			N := 100
			F := 100
			Seed := 123456

			barn := NewBarn(N)
			barn.generateMap(int64(F), int64(Seed))

			barn.Solve()
		}
	})

	b.Run("Solve N1000", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			N := 1000
			F := 100
			Seed := 123456

			barn := NewBarn(N)
			barn.generateMap(int64(F), int64(Seed))

			barn.Solve()
		}
	})
}
