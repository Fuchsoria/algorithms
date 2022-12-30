package hashtable

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTrie(t *testing.T) {
	t.Run("Multi", func(t *testing.T) {
		trie := Constructor()

		trie.Insert("apple")

		found := trie.Search("apple")
		require.True(t, found)

		found = trie.Search("app")
		require.False(t, found)

		trie.Insert("app")

		found = trie.Search("app")
		require.True(t, found)
	})
}
