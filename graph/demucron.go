package graph

import "fmt"

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

func (g *GraphMap) createMatrix() {
	for _, vertex := range g.vertexes {
		for _, subVertex := range g.vertexes {
			if _, ok := g.matrix[vertex]; !ok {
				g.matrix[vertex] = map[Vertex]int{}
			}

			if _, ok := g.matrix[vertex][subVertex]; !ok {
				g.matrix[vertex][subVertex] = 0
			}
		}
	}

	for vertex, edges := range g.data {
		for _, edge := range edges {
			g.matrix[vertex][edge] += 1
		}
	}

	g.generateMatrixResult()
}

func (g *GraphMap) generateMatrixResult() {
	g.vertexesSum = make([]int, len(g.vertexes))

	for _, vertex := range g.vertexes {
		for index, subVertex := range g.vertexes {
			value := g.matrix[vertex][subVertex]

			g.vertexesSum[index] += value
		}
	}
}

func (g *GraphMap) generateLine(vertex Vertex) []int {
	line := []int{}

	for _, subVertex := range g.vertexes {
		value := g.matrix[vertex][subVertex]
		line = append(line, value)
	}

	return line
}

func (g *GraphMap) printMatrix() {
	fmt.Println("-", g.vertexes)

	for _, vertex := range g.vertexes {
		line := g.generateLine(vertex)

		fmt.Println(vertex, line)
	}

	fmt.Println("-", g.vertexesSum)
}

func (g *GraphMap) demucronIteration() []Vertex {
	result := []Vertex{}
	sum := g.vertexesSum

	for sum != nil {
		newSum := sum
		found := false

		for index, value := range sum {
			vertex := g.vertexes[index]

			if value == 0 {
				line := g.generateLine(vertex)

				fmt.Println("MINUS", vertex, line)
				result = append(result, vertex)

				for idx := range newSum {
					newSum[idx] -= line[idx]
				}

				newSum[index] = -1

				found = true
			}
		}

		sum = newSum

		if !found {
			sum = nil
		}
	}

	fmt.Println(result)

	return result
}
