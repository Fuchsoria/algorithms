package graph

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQueue(t *testing.T) {
	t.Run("Multi", func(t *testing.T) {
		queue := NewQueue[string]()
		// A, B, C, D

		queue.Push("A")
		queue.Push("B")
		queue.Push("C")
		queue.Push("D")

		require.Equal(t, queue.Len(), 4)

		a := queue.Shift()
		require.Equal(t, "A", a)
		require.Equal(t, queue.Len(), 3)

		b := queue.Shift()
		require.Equal(t, "B", b)
		require.Equal(t, queue.Len(), 2)

		c := queue.Shift()
		require.Equal(t, "C", c)
		require.Equal(t, queue.Len(), 1)

		d := queue.Shift()
		require.Equal(t, "D", d)
		require.Equal(t, queue.Len(), 0)

		n := queue.Shift()
		require.Equal(t, "", n)
		require.Equal(t, queue.Len(), 0)
	})
}
