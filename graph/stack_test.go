package graph

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStack(t *testing.T) {
	t.Run("Multi", func(t *testing.T) {
		queue := NewStack[string]()
		// A, B, C, D

		queue.Push("A")
		queue.Push("B")
		queue.Push("C")
		queue.Push("D")

		require.Equal(t, queue.Len(), 4)

		d := queue.Pop()
		require.Equal(t, "D", d)
		require.Equal(t, queue.Len(), 3)

		c := queue.Pop()
		require.Equal(t, "C", c)
		require.Equal(t, queue.Len(), 2)

		b := queue.Pop()
		require.Equal(t, "B", b)
		require.Equal(t, queue.Len(), 1)

		a := queue.Pop()
		require.Equal(t, "A", a)
		require.Equal(t, queue.Len(), 0)

		n := queue.Pop()
		require.Equal(t, "", n)
		require.Equal(t, queue.Len(), 0)
	})
}
