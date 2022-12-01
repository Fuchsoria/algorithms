package sort

import (
	"testing"
)

func BenchmarkBubble(b *testing.B) {
	sort := Sort{}

	b.Run("100", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sort.Bubble(case100)
		}
	})

	b.Run("1000", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sort.Bubble(case1000)
		}
	})

	b.Run("10000", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sort.Bubble(case10000)
		}
	})
}

func BenchmarkInsertion(b *testing.B) {
	sort := Sort{}

	b.Run("100", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sort.Insertion(case100)
		}
	})

	b.Run("1000", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sort.Insertion(case1000)
		}
	})

	b.Run("10000", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sort.Insertion(case10000)
		}
	})
}

func BenchmarkInsertionShift(b *testing.B) {
	sort := Sort{}

	b.Run("100", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sort.InsertionShift(case100)
		}
	})

	b.Run("1000", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sort.InsertionShift(case1000)
		}
	})

	b.Run("10000", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sort.InsertionShift(case10000)
		}
	})
}

func BenchmarkInsertionBinary(b *testing.B) {
	sort := Sort{}

	b.Run("100", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sort.InsertionBinarySearch(case100)
		}
	})

	b.Run("1000", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sort.InsertionBinarySearch(case1000)
		}
	})

	b.Run("10000", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sort.InsertionBinarySearch(case10000)
		}
	})
}

func BenchmarkShell(b *testing.B) {
	sort := Sort{}

	b.Run("100", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sort.Shell(case100)
		}
	})

	b.Run("1000", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sort.Shell(case1000)
		}
	})

	b.Run("10000", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sort.Shell(case10000)
		}
	})

	b.Run("100k", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sort.Shell(dataArr100k)
		}
	})

	b.Run("1kk", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sort.Shell(dataArr1kk)
		}
	})
}

func BenchmarkSelection(b *testing.B) {
	sort := Sort{}

	b.Run("100", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sort.Selection(case100)
		}
	})

	b.Run("1000", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sort.Selection(case1000)
		}
	})

	b.Run("10000", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sort.Selection(case10000)
		}
	})

	b.Run("100k", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sort.Selection(dataArr100k)
		}
	})
}

func BenchmarkMerge(b *testing.B) {
	sort := Sort{}

	b.Run("100", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sort.Merge(case100)
		}
	})

	b.Run("1000", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sort.Merge(case1000)
		}
	})

	b.Run("10000", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sort.Merge(case10000)
		}
	})

	b.Run("100k", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sort.Merge(dataArr100k)
		}
	})

	b.Run("1kk", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sort.Merge(dataArr1kk)
		}
	})
}

func BenchmarkQuick(b *testing.B) {
	sort := Sort{}

	b.Run("100", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sort.Quick(case100)
		}
	})

	b.Run("1000", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sort.Quick(case1000)
		}
	})

	b.Run("10000", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sort.Quick(case10000)
		}
	})

	b.Run("100k", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sort.Quick(dataArr100k)
		}
	})

	b.Run("1kk", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sort.Quick(dataArr1kk)
		}
	})
}

func BenchmarkBucket(b *testing.B) {
	sort := SortL{}

	b.Run("100", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sort.Bucket(case100, 65535)
		}
	})

	b.Run("1000", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sort.Bucket(case1000, 65535)
		}
	})

	b.Run("10000", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sort.Bucket(case10000, 65535)
		}
	})

	b.Run("100k", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sort.Bucket(dataArr100k, 65535)
		}
	})

	b.Run("1kk", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sort.Bucket(dataArr1kk, 65535)
		}
	})
}

func BenchmarkCounting(b *testing.B) {
	sort := SortL{}

	b.Run("100", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sort.Counting(case100, 65535)
		}
	})

	b.Run("1000", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sort.Counting(case1000, 65535)
		}
	})

	b.Run("10000", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sort.Counting(case10000, 65535)
		}
	})

	b.Run("100k", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sort.Counting(dataArr100k, 65535)
		}
	})

	b.Run("1kk", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sort.Counting(dataArr1kk, 65535)
		}
	})
}

func BenchmarkRadix(b *testing.B) {
	sort := SortL{}

	b.Run("100", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sort.Radix(case100)
		}
	})

	b.Run("1000", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sort.Radix(case1000)
		}
	})

	b.Run("10000", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sort.Radix(case10000)
		}
	})

	b.Run("100k", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sort.Radix(dataArr100k)
		}
	})

	b.Run("1kk", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			sort.Radix(dataArr1kk)
		}
	})
}
