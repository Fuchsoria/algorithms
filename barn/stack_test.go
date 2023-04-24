package barn

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStack(t *testing.T) {
	t.Run("Multi", func(t *testing.T) {
		stack := NewStack[string]()
		// A, B, C, D

		stack.Push("A")
		stack.Push("B")
		stack.Push("C")
		stack.Push("D")

		require.Equal(t, stack.Len(), 4)

		d := stack.Pop()
		require.Equal(t, "D", d)
		require.Equal(t, stack.Len(), 3)

		c := stack.Pop()
		require.Equal(t, "C", c)
		require.Equal(t, stack.Len(), 2)

		b := stack.Pop()
		require.Equal(t, "B", b)
		require.Equal(t, stack.Len(), 1)

		a := stack.Pop()
		require.Equal(t, "A", a)
		require.Equal(t, stack.Len(), 0)

		n := stack.Pop()
		require.Equal(t, "", n)
		require.Equal(t, stack.Len(), 0)
	})
}
