package graph

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestKosaraju(t *testing.T) {
	t.Run("Multi", func(t *testing.T) {
		expected := [][]int{{0, 1, 2}, {3}, {4}}

		g := New(5)
		g.addEdge(1, 0)
		g.addEdge(0, 2)
		g.addEdge(2, 1)
		g.addEdge(0, 3)
		g.addEdge(3, 4)

		require.Equal(t, expected, g.printSCCs())
	})
}
