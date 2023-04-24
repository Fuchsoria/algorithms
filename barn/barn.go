package barn

import (
	"fmt"
	"math/rand"
)

type Barn struct {
	N    int
	m    [][]bool
	line []int
	L    []int
	R    []int
}

func (b *Barn) generateMap(F, seed int64) {
	random := rand.New(rand.NewSource(seed))

	for y := 0; y < b.N; y++ {
		for x := 0; x < b.N; x++ {
			b.m[x][y] = random.Int63n(F) == 0
		}
	}
}

func (b *Barn) calculateLineHeights(y int) {
	for x := 0; x < b.N; x++ {
		if b.m[x][y] {
			b.line[x] = 0
		} else {
			b.line[x]++
		}
	}
}

func (b *Barn) calculateLeftRanges(y int) {
	stack := NewStack[int]()

	for x := b.N - 1; x >= 0; x-- {
		for stack.Len() > 0 {
			if b.line[stack.list.Back().Value.(int)] > b.line[x] {
				b.L[stack.Pop()] = x + 1
			} else {
				break
			}
		}

		stack.Push(x)
	}

	for stack.Len() > 0 {
		b.L[stack.Pop()] = 0
	}
}

func (b *Barn) calculateRightRanges(y int) {
	stack := NewStack[int]()

	for x := 0; x < b.N; x++ {
		for stack.Len() > 0 {
			if b.line[stack.list.Back().Value.(int)] > b.line[x] {
				b.R[stack.Pop()] = x - 1
			} else {
				break
			}
		}

		stack.Push(x)
	}

	for stack.Len() > 0 {
		b.R[stack.Pop()] = b.N - 1
	}
}

func (b *Barn) Solve() int {
	maxSquare := 0

	for y := 0; y < b.N; y++ {
		b.calculateLineHeights(y)
		b.calculateRightRanges(y)
		b.calculateLeftRanges(y)

		for x := 0; x < b.N; x++ {
			height := b.line[x]
			width := b.R[x] - b.L[x] + 1
			square := height * width

			if maxSquare < square {
				maxSquare = square
			}
		}
	}

	return maxSquare
}

func (b *Barn) Print() {
	if b.N > 12 {
		return
	}

	for y := 0; y < b.N; y++ {
		for x := 0; x < b.N; x++ {
			if b.m[x][y] {
				fmt.Print("#")
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Println()
	}
}

func NewBarn(N int) *Barn {
	arr := make([][]bool, N)
	for i := range arr {
		arr[i] = make([]bool, N)
	}

	return &Barn{N: N, m: arr, line: make([]int, N), L: make([]int, N), R: make([]int, N)}
}
