package graph

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDemucron(t *testing.T) {
	t.Run("Multi", func(t *testing.T) {
		expected := []Vertex{"C", "D", "F", "A", "B", "E", "G", "H"}

		g := NewGraphMap()

		g.addVertex("A")
		g.addVertex("B")
		g.addVertex("C")
		g.addVertex("D")
		g.addVertex("E")
		g.addVertex("F")
		g.addVertex("G")
		g.addVertex("H")

		g.addEdge("A", "B")
		g.addEdge("B", "E")
		g.addEdge("C", "D")
		g.addEdge("D", "A")
		g.addEdge("D", "B")
		g.addEdge("D", "E")
		g.addEdge("D", "F")
		g.addEdge("E", "G")
		g.addEdge("F", "E")
		g.addEdge("F", "H")
		g.addEdge("G", "H")

		g.createMatrix()

		g.printMatrix()

		result := g.demucronIteration()
		require.Equal(t, expected, result)
	})
}
