package priorityqueue

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQueue(t *testing.T) {
	t.Run("add", func(t *testing.T) {
		queue := Queue{}

		queue.Add(10)
		queue.Add(20)
		queue.Add(30)

		require.Equal(t, []int64{10, 20, 30}, queue.GetValues())
	})

	t.Run("take", func(t *testing.T) {
		queue := Queue{}

		queue.Add(10)
		queue.Add(20)
		node := queue.Take()
		queue.Add(30)

		require.NotNil(t, node)
		require.Equal(t, int64(10), node.Value)
		require.Equal(t, []int64{20, 30}, queue.GetValues())
	})
}
