package compression

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCompression(t *testing.T) {
	t.Run("String Multi", func(t *testing.T) {
		input := "wwwwaaadexxxxxxywww"
		output := "4w3a1d1e6x1y3w"

		c := NewCompression()

		encoded := c.Encode(input)
		require.Equal(t, output, encoded)

		decoded := c.Decode(encoded)
		require.Equal(t, input, decoded)
	})
}
