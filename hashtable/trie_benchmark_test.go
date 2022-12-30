package hashtable

import (
	"fmt"
	"testing"
)

func BenchmarkTrie(b *testing.B) {
	b.Run("insert", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			trie := Constructor()

			for i := 0; i < 10000; i++ {
				trie.Insert(fmt.Sprint(i) + "_value")
			}
		}
	})

	b.Run("insert & search", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			trie := Constructor()

			for i := 0; i < 10000; i++ {
				trie.Insert(fmt.Sprint(i) + "_value")
			}

			for i := 0; i < 10000; i++ {
				trie.Search(fmt.Sprint(i) + "_value")
			}
		}
	})
}
