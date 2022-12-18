package hashtable

import (
	"fmt"
	"testing"
)

func BenchmarkHashtable(b *testing.B) {
	b.Run("insert", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			hashtable := New(10)

			for i := 0; i < 10000; i++ {
				hashtable.Insert(fmt.Sprint(i), fmt.Sprint(i)+"_value")
			}
		}
	})

	b.Run("insert & get", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			hashtable := New(10)

			for i := 0; i < 10000; i++ {
				hashtable.Insert(fmt.Sprint(i), fmt.Sprint(i)+"_value")
			}

			for i := 0; i < 10000; i++ {
				hashtable.Get(fmt.Sprint(i))
			}
		}
	})

	b.Run("insert & delete", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			hashtable := New(10)

			for i := 0; i < 10000; i++ {
				hashtable.Insert(fmt.Sprint(i), fmt.Sprint(i)+"_value")
			}

			for i := 0; i < 10000; i++ {
				hashtable.Delete(fmt.Sprint(i))
			}
		}
	})
}
