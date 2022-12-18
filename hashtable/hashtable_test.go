package hashtable

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHashTable(t *testing.T) {
	t.Run("Multi", func(t *testing.T) {
		input := []string{"First", "Second", "Third", "Fourth", "Fifth", "Sixth", "Seventh", "Eighth", "Ninth", "Tenth"}

		hashtable := New(10)

		for i, k := range input {
			hashtable.Insert(k, fmt.Sprintf("%s-%d", k, i+1))
		}

		for i, k := range input {
			v, ok := hashtable.Get(k)

			require.True(t, ok)
			require.Equal(t, v, fmt.Sprintf("%s-%d", k, i+1))
		}

		for _, k := range input[5:] {
			err := hashtable.Delete(k)

			require.Nil(t, err)
		}

		for i, k := range input[:5] {
			v, ok := hashtable.Get(k)

			require.True(t, ok)
			require.Equal(t, v, fmt.Sprintf("%s-%d", k, i+1))
		}

		for _, k := range input[5:] {
			_, ok := hashtable.Get(k)

			require.False(t, ok)
		}
	})
}
