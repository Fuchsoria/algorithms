package graph

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestKruskals(t *testing.T) {
	t.Run("Multi", func(t *testing.T) {
		expected := []KruskalResultItem{
			{0, 0, 1, 2},
			{1, 1, 2, 3},
			{2, 1, 4, 5},
			{3, 0, 3, 6},
		}

		cost := [][]int{
			{INF, 2, INF, 6, INF},
			{2, INF, 3, 8, 5},
			{INF, 3, INF, INF, 7},
			{6, 8, INF, INF, 9},
			{INF, 5, 7, 9, INF},
		}

		k := NewKruskal(5)
		results, min := k.MST(cost)

		require.Equal(t, expected, results)
		require.Equal(t, 16, min)
	})
}
