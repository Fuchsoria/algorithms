package graph

type DijkstrasResultItem struct {
	vertex      int
	distFromSrc int
}

type Dijkstras struct {
	V int
}

func (d *Dijkstras) minDistance(dist []int, sptSet []bool) int {
	min := INF
	minIndex := -1

	for v := 0; v < d.V; v++ {
		if !sptSet[v] && dist[v] <= min {
			min = dist[v]
			minIndex = v
		}
	}

	return minIndex
}

func (d *Dijkstras) result(dist []int, src int) ([]DijkstrasResultItem, DijkstrasResultItem) {
	result := []DijkstrasResultItem{}
	minDist := DijkstrasResultItem{0, INF}

	for i := 0; i < d.V; i++ {
		result = append(result, DijkstrasResultItem{i, dist[i]})

		if src != i && dist[i] < minDist.distFromSrc {
			minDist = DijkstrasResultItem{i, dist[i]}
		}
	}

	return result, minDist
}

func (d *Dijkstras) Dijkstra(graph [][]int, src int) ([]DijkstrasResultItem, DijkstrasResultItem) {
	dist := make([]int, d.V)
	sptSet := make([]bool, d.V)

	for i := 0; i < d.V; i++ {
		dist[i] = INF
		sptSet[i] = false
	}

	dist[src] = 0

	for count := 0; count < d.V-1; count++ {
		u := d.minDistance(dist, sptSet)

		sptSet[u] = true

		for v := 0; v < d.V; v++ {
			if !sptSet[v] && graph[u][v] != 0 && dist[u] != INF && dist[u]+graph[u][v] < dist[v] {
				dist[v] = dist[u] + graph[u][v]
			}
		}
	}

	return d.result(dist, src)
}

func NewDijkstras(v int) *Dijkstras {
	return &Dijkstras{V: v}
}
