package bellman

import (
	"errors"
)

type Algorithm struct {
	DistanceMap map[int]int
	size        int
	EdgeList    []*Edge
}

type Edge struct {
	from   int
	to     int
	weight int
}

var (
	ErrUnevenMatrix = errors.New("adjacency matrix is uneven")
	InfVal          = 16
)

// New initializes a new run of the Bellman ford algorithm
func (a *Algorithm) New(graph [][]int) error {
	a.size = len(graph)
	expectedTotal := a.size * a.size
	cellTotal := 0
	for _, h := range graph {
		cellTotal = +len(h)
	}
	if cellTotal != expectedTotal {
		return ErrUnevenMatrix
	}

	// create new distance map
	for i := 0; i < a.size; i++ {
		a.DistanceMap[i] = InfVal
	}

	// generate edgelist
	for i := 0; i < a.size; i++ {
		for j := 0; j < a.size; j++ {
			if i != j && graph[i][j] > 0 {
				a.EdgeList = append(a.EdgeList, &Edge{from: i, to: i, weight: graph[i][j]})
			}
		}
	}

	return nil
}

// Run starts running the algorithm
