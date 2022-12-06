package tree

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBST(t *testing.T) {
	t.Run("Multi", func(t *testing.T) {
		input := []int{8, 9, 6, 7, 3, 4, 2, 1, 5, 0}
		expectInsert := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		expectRemove := []int{0, 1, 2, 3, 4, 6, 7, 8, 9}

		bst := bstTree{}

		for _, v := range input {
			bst.Insert(v)
		}

		require.Equal(t, expectInsert, bst.Inorder())

		node, err := bst.Search(5)
		require.NoError(t, err)
		require.NotNil(t, node)
		require.Equal(t, 5, node.Value)

		bst.Remove(5)

		node, err = bst.Search(5)
		require.ErrorIs(t, ErrNodeNotFound, err)
		require.Nil(t, node)

		require.Equal(t, expectRemove, bst.Inorder())
	})
}
