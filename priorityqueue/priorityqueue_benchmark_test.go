package priorityqueue

import (
	"math/rand"
	"testing"
)

func BenchmarkPriorityQueue(b *testing.B) {
	for _, bm := range Cases {
		b.Run(bm.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				queue := PriorityQueue{}

				for j := 0; j < bm.Input; j++ {
					queue.Enqueue(rand.Int63n(100000), int64(j))
				}
			}
		})
	}
}

func BenchmarkPriorityQueueMapped(b *testing.B) {
	for _, bm := range Cases {
		b.Run(bm.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				queue := PriorityQueueMapped{}

				for j := 0; j < bm.Input; j++ {
					queue.Enqueue(rand.Int63n(100000), int64(j))
				}
			}
		})
	}
}
