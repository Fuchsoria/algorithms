package graph

import (
	"math"
)

var INF = math.MaxInt

type KruskalResultItem struct {
	edgeCount int
	a         int
	b         int
	min       int
}

type Kruskal struct {
	V      int
	parent []int
}

func (k *Kruskal) find(i int) int {
	for k.parent[i] != i {
		i = k.parent[i]
	}

	return i
}

func (k *Kruskal) union(i, j int) {
	a := k.find(i)
	b := k.find(j)

	k.parent[a] = b
}

func (k *Kruskal) MST(cost [][]int) ([]KruskalResultItem, int) {
	minCost := 0
	resultEdges := []KruskalResultItem{}

	for i := 0; i < k.V; i++ {
		k.parent[i] = i
	}

	edgeCount := 0
	for edgeCount < k.V-1 {
		min := INF
		a := -1
		b := -1

		for i := 0; i < k.V; i++ {
			for j := 0; j < k.V; j++ {
				if k.find(i) != k.find(j) && cost[i][j] < min {
					min = cost[i][j]
					a = i
					b = j
				}
			}
		}

		k.union(a, b)

		resultEdges = append(resultEdges, KruskalResultItem{edgeCount, a, b, min})

		edgeCount += 1
		minCost += min
	}

	return resultEdges, minCost
}

func NewKruskal(v int) *Kruskal {
	return &Kruskal{
		V:      v,
		parent: make([]int, v),
	}
}
