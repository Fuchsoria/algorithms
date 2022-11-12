package priorityqueue

import "testing"

func BenchmarkQueue(b *testing.B) {
	for _, bm := range Cases {
		b.Run(bm.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				queue := Queue{}

				for j := 0; j < bm.Input; j++ {
					queue.Add(int64(j))
				}
			}
		})
	}
}
