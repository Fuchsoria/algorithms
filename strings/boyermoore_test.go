package strs

import (
	"testing"

	"github.com/stretchr/testify/require"
	//	"github.com/stretchr/testify/require"
)

func TestBoyerMoore(t *testing.T) {
	t.Run("Multi", func(t *testing.T) {
		full := "ABAAABCDLHEREABCD"
		part := "ABC"
		expected := []int{4, 13}
		expectedIndex := 4

		bm := NewBM()
		bm2 := NewBMUpdated()

		has, startIndexes := bm.search(full, part)

		require.True(t, has)
		require.Equal(t, expected, startIndexes)

		index := bm2.Find(full, part)
		require.Equal(t, expectedIndex, index)

		index = bm2.Find2(full, part)
		require.Equal(t, expectedIndex, index)

		index = bm2.Find3(full, part)
		require.Equal(t, expectedIndex, index)
	})
}
