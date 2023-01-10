package graph

type (
	Vertex string
	Edges  []Vertex
)

type GraphMap struct {
	vertexes    []Vertex
	vertexesSum []int
	data        map[Vertex]Edges
	matrix      map[Vertex]map[Vertex]int
}

func NewGraphMap() *GraphMap {
	return &GraphMap{
		data:   map[Vertex]Edges{},
		matrix: map[Vertex]map[Vertex]int{},
	}
}

func (g *GraphMap) DFS(edge Vertex) []Vertex {
	result := []Vertex{}
	visited := map[Vertex]bool{}
	stack := NewStack[Vertex]()

	stack.Push(edge)

	for stack.Len() > 0 {
		item := stack.Pop()

		if !visited[item] {
			result = append(result, item)
			visited[item] = true
		}

		for _, v := range g.data[item] {
			if !visited[v] {
				stack.Push(v)
			}
		}
	}

	return result
}

func (g *GraphMap) BFS(edge Vertex) []Vertex {
	result := []Vertex{}
	visited := map[Vertex]bool{}
	queue := NewQueue[Vertex]()

	queue.Push(edge)

	for queue.Len() > 0 {
		item := queue.Shift()

		if !visited[item] {
			result = append(result, item)
			visited[item] = true
		}

		for _, v := range g.data[item] {
			if !visited[v] {
				queue.Push(v)
			}
		}
	}

	return result
}

func (g *GraphMap) addVertex(v Vertex) {
	if _, ok := g.data[v]; !ok {
		g.data[v] = Edges{}
		g.vertexes = append(g.vertexes, v)
	}
}

func (g *GraphMap) addEdge(v Vertex, s Vertex) {
	if _, ok := g.data[v]; !ok {
		g.addVertex(v)
	}

	if _, ok := g.data[s]; !ok {
		g.addVertex(s)
	}

	if edges, ok := g.data[v]; ok {
		g.data[v] = append(edges, s)
	}
}
