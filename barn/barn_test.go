package barn

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBarn(t *testing.T) {
	t.Run("Multitest", func(t *testing.T) {
		N := 12
		F := 30
		Seed := 12345

		barn := NewBarn(N)
		barn.generateMap(int64(F), int64(Seed))
		barn.Print()

		maxSquare := barn.Solve()

		require.Equal(t, 54, maxSquare)
	})
}
