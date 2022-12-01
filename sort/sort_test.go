package sort

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBubble(t *testing.T) {
	sort := Sort{}

	t.Run("sort", func(t *testing.T) {
		input := []int{8, 9, 6, 7, 3, 4, 2, 1, 5, 0}
		result := sort.Bubble(input)

		require.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, result)
	})
}

func TestInsertion(t *testing.T) {
	sort := Sort{}

	t.Run("sort", func(t *testing.T) {
		input := []int{8, 9, 6, 7, 3, 4, 2, 1, 5, 0}
		result := sort.Insertion(input)

		require.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, result)
	})
}

func TestInsertionShift(t *testing.T) {
	sort := Sort{}

	t.Run("sort", func(t *testing.T) {
		input := []int{8, 9, 6, 7, 3, 4, 2, 1, 5, 0}
		result := sort.InsertionShift(input)

		require.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, result)
	})
}

func TestInsertionBinarySearch(t *testing.T) {
	sort := Sort{}

	t.Run("sort", func(t *testing.T) {
		input := []int{8, 9, 6, 7, 3, 4, 2, 1, 5, 0}
		result := sort.InsertionBinarySearch(input)

		require.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, result)
	})
}

func TestShell(t *testing.T) {
	sort := Sort{}

	t.Run("sort", func(t *testing.T) {
		input := []int{8, 9, 6, 7, 3, 4, 2, 1, 5, 0}
		result := sort.Shell(input)

		require.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, result)
	})
}

func TestSelection(t *testing.T) {
	sort := Sort{}

	t.Run("sort", func(t *testing.T) {
		input := []int{8, 9, 6, 7, 3, 4, 2, 1, 5, 0}
		result := sort.Selection(input)

		require.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, result)
	})
}

func TestHeap(t *testing.T) {
	sort := Sort{}

	t.Run("sort", func(t *testing.T) {
		input := []int{8, 9, 6, 7, 3, 4, 2, 1, 5, 0}
		result := sort.Heap(input)

		require.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, result)
	})
}

func TestQuick(t *testing.T) {
	sort := Sort{}

	t.Run("sort", func(t *testing.T) {
		input := []int{8, 9, 6, 7, 3, 4, 2, 1, 5, 0}
		result := sort.Quick(input)

		require.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, result)
	})
}

func TestMerge(t *testing.T) {
	sort := Sort{}

	t.Run("sort", func(t *testing.T) {
		input := []int{8, 9, 6, 7, 3, 4, 2, 1, 5, 0}
		result := sort.Merge(input)

		require.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, result)
	})
}

func TestBucket(t *testing.T) {
	sort := SortL{}

	t.Run("sort", func(t *testing.T) {
		input := []int{8, 9, 6, 7, 3, 4, 2, 1, 5, 0}
		result := sort.Bucket(input, 10)

		require.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, result)
	})
}

func TestCounting(t *testing.T) {
	sort := SortL{}

	t.Run("sort", func(t *testing.T) {
		input := []int{8, 9, 6, 7, 3, 4, 2, 1, 5, 0}
		result := sort.Counting(input, 10)

		require.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, result)
	})
}

func TestRadix(t *testing.T) {
	sort := SortL{}

	t.Run("sort", func(t *testing.T) {
		input := []int{8, 9, 6, 7, 3, 4, 2, 1, 5, 0}
		result := sort.Radix(input)

		require.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, result)
	})
}
