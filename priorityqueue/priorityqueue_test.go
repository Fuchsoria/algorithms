package priorityqueue

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPriorityQueue(t *testing.T) {
	t.Run("combo", func(t *testing.T) {
		queue := PriorityQueue{}

		queue.Enqueue(1, 10)
		queue.Enqueue(2, 20)
		queue.Enqueue(3, 30)
		queue.Enqueue(4, 40)
		queue.Enqueue(5, 50)
		queue.Enqueue(5, 55)
		queue.Enqueue(5, 56)
		queue.Enqueue(6, 60)
		queue.Enqueue(7, 70)
		queue.Enqueue(8, 80)

		require.Equal(t,
			"[{1 [10]} {2 [20]} {3 [30]} {4 [40]} {5 [50 55 56]} {6 [60]} {7 [70]} {8 [80]}]",
			fmt.Sprint(queue.GetPrioritizedValues()),
		)

		value, err := queue.Dequeue()
		require.NoError(t, err)
		require.Equal(t, int64(80), value)

		require.Equal(t,
			"[{1 [10]} {2 [20]} {3 [30]} {4 [40]} {5 [50 55 56]} {6 [60]} {7 [70]}]",
			fmt.Sprint(queue.GetPrioritizedValues()),
		)

		value, err = queue.Dequeue()
		require.NoError(t, err)
		require.Equal(t, int64(70), value)

		require.Equal(t,
			"[{1 [10]} {2 [20]} {3 [30]} {4 [40]} {5 [50 55 56]} {6 [60]}]",
			fmt.Sprint(queue.GetPrioritizedValues()),
		)

		value, err = queue.Dequeue()
		require.NoError(t, err)
		require.Equal(t, int64(60), value)

		require.Equal(t,
			"[{1 [10]} {2 [20]} {3 [30]} {4 [40]} {5 [50 55 56]}]",
			fmt.Sprint(queue.GetPrioritizedValues()),
		)

		value, err = queue.Dequeue()
		require.NoError(t, err)
		require.Equal(t, int64(50), value)

		require.Equal(t,
			"[{1 [10]} {2 [20]} {3 [30]} {4 [40]} {5 [55 56]}]",
			fmt.Sprint(queue.GetPrioritizedValues()),
		)

		require.Equal(t, []int64{1, 2, 3, 4, 5}, queue.GetPriorities())
	})

	t.Run("empty", func(t *testing.T) {
		queue := PriorityQueue{}

		queue.Enqueue(1, 10)

		require.Equal(t,
			"[{1 [10]}]",
			fmt.Sprint(queue.GetPrioritizedValues()),
		)

		value, err := queue.Dequeue()
		require.NoError(t, err)
		require.Equal(t, int64(10), value)

		require.Equal(t, []int64{}, queue.GetPriorities())

		value, err = queue.Dequeue()
		require.ErrorIs(t, err, errorEmpty)
		require.Equal(t, int64(0), value)

		require.Equal(t, []int64{}, queue.GetPriorities())
	})
}
