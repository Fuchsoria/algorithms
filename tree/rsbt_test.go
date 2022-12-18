package tree

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRBST(t *testing.T) {
	t.Run("Multi", func(t *testing.T) {
		input := []int{8, 9, 6, 7, 3, 4, 2, 1, 5, 0}
		expectInsert := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		expectRemove := []int{0, 1, 2, 3, 4, 6, 7, 8, 9}

		rbst := rbsTree{}

		for _, v := range input {
			rbst.Insert(v)
		}

		require.Equal(t, expectInsert, rbst.Inorder())

		node, err := rbst.Search(5)
		require.NoError(t, err)
		require.NotNil(t, node)
		require.Equal(t, 5, node.Value)

		rbst.Remove(5)

		node, err = rbst.Search(5)
		require.ErrorIs(t, ErrNodeNotFound, err)
		require.Nil(t, node)

		require.Equal(t, expectRemove, rbst.Inorder())
	})
}
