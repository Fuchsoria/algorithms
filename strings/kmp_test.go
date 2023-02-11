package strs

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestKMP(t *testing.T) {
	t.Run("Multi", func(t *testing.T) {
		full := "ABAAABCDLHEREABCD"
		part := "ABC"
		expected := []int{4, 13}

		kmp := NewKMP()

		has, startIndexes := kmp.Search(full, part)

		require.True(t, has)
		require.Equal(t, expected, startIndexes)
	})
}
