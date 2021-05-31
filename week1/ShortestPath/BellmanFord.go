package ShortestPath

import (
	"../Graph"
	"math"
)

type BellmanFord struct {
	SourceVertex int
	G            Graph.Graph
	A            [][]float64
}

func InitBellmanFord(g Graph.Graph, sourceVertex int) BellmanFord {
	b := BellmanFord{
		SourceVertex: sourceVertex,
		G:            g,
		A:            make([][]float64, g.NumVertices),
	}

	for numEdgeBudget := 0; numEdgeBudget < len(b.A); numEdgeBudget++ {
		b.A[numEdgeBudget] = make([]float64, g.NumVertices)
	}

	for nodeIndx := 0; nodeIndx < g.NumVertices; nodeIndx++ {
		b.A[0][nodeIndx] = math.Inf(1)
	}

	b.A[0][sourceVertex-1] = 0

	return b
}
