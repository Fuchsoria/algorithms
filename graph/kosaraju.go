package graph

import (
	"container/list"
)

type Graph struct {
	V   int
	adj []*list.List
}

func (g *Graph) addEdge(v, w int) {
	g.adj[v].PushBack(w)
}

func (g *Graph) DFSUtil(v int, visited []bool, fn func(v int)) {
	visited[v] = true
	fn(v)

	l := g.adj[v]

	for e := l.Front(); e != nil; e = e.Next() {
		val := e.Value.(int)

		if !visited[val] {
			g.DFSUtil(val, visited, fn)
		}
	}
}

func (g *Graph) getTranspose() *Graph {
	newGraph := New(g.V)

	for v := 0; v < g.V; v++ {
		l := g.adj[v]

		for e := l.Front(); e != nil; e = e.Next() {
			val := e.Value.(int)

			newGraph.addEdge(val, v)
		}
	}

	return newGraph
}

func (g *Graph) fillOrder(v int, visited []bool, stack *list.List) {
	visited[v] = true

	l := g.adj[v]

	for e := l.Front(); e != nil; e = e.Next() {
		val := e.Value.(int)

		if !visited[val] {
			g.fillOrder(val, visited, stack)
		}
	}

	stack.PushBack(v)
}

func (g *Graph) printSCCs() [][]int {
	stack := list.New()
	visited := make([]bool, g.V)
	result := [][]int{}

	for i := 0; i < g.V; i++ {
		if !visited[i] {
			g.fillOrder(i, visited, stack)
		}
	}

	gr := g.getTranspose()

	for i := 0; i < g.V; i++ {
		visited[i] = false
	}

	for e := stack.Back(); e != nil; e = e.Prev() {
		val := e.Value.(int)

		if !visited[val] {
			subResult := []int{}

			gr.DFSUtil(val, visited, func(v int) { subResult = append(subResult, v) })

			result = append(result, subResult)
		}
	}

	return result
}

func New(v int) *Graph {
	g := &Graph{}
	g.V = v
	g.adj = make([]*list.List, v)

	for i := 0; i < v; i++ {
		g.adj[i] = list.New()
	}

	return g
}
