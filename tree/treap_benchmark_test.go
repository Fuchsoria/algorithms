package tree

import (
	"testing"
)

func BenchmarkTreap(b *testing.B) {
	b.Run("insert", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			tree := TreapTree{}

			for _, v := range dataArr10k {
				tree.Insert(v)
			}
		}
	})

	b.Run("insert & search", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			tree := TreapTree{}

			for _, v := range dataArr10k {
				tree.Insert(v)
			}

			for _, v := range dataArr1k {
				tree.Search(v)
			}
		}
	})

	b.Run("insert & remove", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			tree := TreapTree{}

			for _, v := range dataArr10k {
				tree.Insert(v)
			}

			for _, v := range dataArr1k {
				tree.Remove(v)
			}
		}
	})

	b.Run("insert & print", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			tree := TreapTree{}

			for _, v := range dataArr10k {
				tree.Insert(v)
			}

			tree.Inorder()
		}
	})
}

func BenchmarkSortedTreap(b *testing.B) {
	b.Run("insert", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			tree := TreapTree{}

			for _, v := range dataArrSorted10k {
				tree.Insert(v)
			}
		}
	})

	b.Run("insert & search", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			tree := TreapTree{}

			for _, v := range dataArrSorted10k {
				tree.Insert(v)
			}

			for _, v := range dataArrSorted1k {
				tree.Search(v)
			}
		}
	})

	b.Run("insert & remove", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			tree := TreapTree{}

			for _, v := range dataArrSorted10k {
				tree.Insert(v)
			}

			for _, v := range dataArrSorted1k {
				tree.Remove(v)
			}
		}
	})

	b.Run("insert & print", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			tree := TreapTree{}

			for _, v := range dataArrSorted10k {
				tree.Insert(v)
			}

			tree.Inorder()
		}
	})
}
