package arrays

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSingleArray(t *testing.T) {
	t.Run("add", func(t *testing.T) {
		inst := NewSingleArr[int]()

		inst.Add(10)
		inst.Add(11)
		inst.Add(12)

		require.Equal(t, []int{10, 11, 12}, inst.GetArray())
	})

	t.Run("add by index", func(t *testing.T) {
		inst := NewSingleArr[int]()

		inst.Add(10)
		inst.Add(11)
		inst.Add(12)
		inst.AddByIndex(14, 2)

		require.Equal(t, []int{10, 11, 14, 12}, inst.GetArray())
	})

	t.Run("remove", func(t *testing.T) {
		inst := NewSingleArr[int]()

		inst.Add(10)
		inst.Add(11)
		inst.Add(12)
		inst.AddByIndex(14, 2)

		result := inst.Remove(2)

		require.Equal(t, []int{10, 11, 12}, inst.GetArray())
		require.Equal(t, 14, result)
	})

	t.Run("get", func(t *testing.T) {
		inst := NewSingleArr[int]()

		inst.Add(10)
		inst.Add(11)
		inst.Add(12)
		inst.AddByIndex(14, 2)

		result := inst.Get(2)

		require.Equal(t, 14, result)
	})
}

func TestVectorArray(t *testing.T) {
	t.Run("add", func(t *testing.T) {
		inst := NewVectorArr(10)

		inst.Add(10)
		inst.Add(11)
		inst.Add(12)

		require.Equal(t, []int{10, 11, 12}, inst.GetArray())
	})

	t.Run("add by index", func(t *testing.T) {
		inst := NewVectorArr(3)

		inst.Add(10)
		inst.Add(11)
		inst.Add(12)
		inst.AddByIndex(14, 2)

		require.Equal(t, []int{10, 11, 14, 12}, inst.GetArray())
	})

	t.Run("remove", func(t *testing.T) {
		inst := NewVectorArr(3)

		inst.Add(10)
		inst.Add(11)
		inst.Add(12)
		inst.AddByIndex(14, 2)

		result := inst.Remove(2)

		require.Equal(t, []int{10, 11, 12}, inst.GetArray())
		require.Equal(t, 14, result)
	})

	t.Run("get", func(t *testing.T) {
		inst := NewVectorArr(3)

		inst.Add(10)
		inst.Add(11)
		inst.Add(12)
		inst.AddByIndex(14, 2)

		result := inst.Get(2)

		require.Equal(t, 14, result)
	})
}

func TestFactorArray(t *testing.T) {
	t.Run("add", func(t *testing.T) {
		inst := NewFactorArr()

		inst.Add(10)
		inst.Add(11)
		inst.Add(12)

		require.Equal(t, []int{10, 11, 12}, inst.GetArray())
	})

	t.Run("add by index", func(t *testing.T) {
		inst := NewFactorArr()

		inst.Add(10)
		inst.Add(11)
		inst.Add(12)
		inst.AddByIndex(14, 2)

		require.Equal(t, []int{10, 11, 14, 12}, inst.GetArray())
	})

	t.Run("remove", func(t *testing.T) {
		inst := NewFactorArr()

		inst.Add(10)
		inst.Add(11)
		inst.Add(12)
		inst.AddByIndex(14, 2)

		result := inst.Remove(2)

		require.Equal(t, []int{10, 11, 12}, inst.GetArray())
		require.Equal(t, 14, result)
	})

	t.Run("get", func(t *testing.T) {
		inst := NewFactorArr()

		inst.Add(10)
		inst.Add(11)
		inst.Add(12)
		inst.AddByIndex(14, 2)

		result := inst.Get(2)

		require.Equal(t, 14, result)
	})
}

func TestArrayList(t *testing.T) {
	t.Run("add", func(t *testing.T) {
		inst := NewArrayList[int]()

		inst.Add(10)
		inst.Add(11)
		inst.Add(12)

		require.Equal(t, []int{10, 11, 12}, inst.GetArray())
	})

	t.Run("add by index", func(t *testing.T) {
		inst := NewFactorArr()

		inst.Add(10)
		inst.Add(11)
		inst.Add(12)
		inst.AddByIndex(14, 2)

		require.Equal(t, []int{10, 11, 14, 12}, inst.GetArray())
	})

	t.Run("remove", func(t *testing.T) {
		inst := NewArrayList[int]()

		inst.Add(10)
		inst.Add(11)
		inst.Add(12)
		inst.AddByIndex(14, 2)

		result := inst.Remove(2)

		require.Equal(t, []int{10, 11, 12}, inst.GetArray())
		require.Equal(t, 14, result)
	})

	t.Run("get", func(t *testing.T) {
		inst := NewArrayList[int]()

		inst.Add(10)
		inst.Add(11)
		inst.Add(12)
		inst.AddByIndex(14, 2)

		result := inst.Get(2)

		require.Equal(t, 14, result)
	})
}
