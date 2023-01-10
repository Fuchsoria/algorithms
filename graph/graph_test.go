package graph

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSearch(t *testing.T) {
	t.Run("DFS & BFS", func(t *testing.T) {
		g := NewGraphMap()

		g.addVertex("A")
		g.addVertex("B")
		g.addVertex("C")
		g.addVertex("D")

		g.addEdge("A", "B")
		g.addEdge("A", "C")
		g.addEdge("B", "C")
		g.addEdge("C", "A")
		g.addEdge("C", "D")
		g.addEdge("D", "D")

		resultDFS := g.DFS("A")
		expectedDFS := []Vertex{"A", "C", "D", "B"}
		require.Equal(t, expectedDFS, resultDFS)

		expectedBFS := []Vertex{"A", "B", "C", "D"}
		resultBFS := g.BFS("A")
		require.Equal(t, expectedBFS, resultBFS)
	})
}
