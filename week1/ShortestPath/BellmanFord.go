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

func (b *BellmanFord) Evaluate() {
	for numEdgeBudget := 1; numEdgeBudget <= b.G.NumVertices; numEdgeBudget++ {
		for nodeIdx := 0; nodeIdx < b.G.NumVertices; nodeIdx++ {
			b.A[numEdgeBudget][nodeIdx] = b.A[numEdgeBudget-1][nodeIdx]
			minOfIncomingToNodeDest := float64(0)
			nodeDest := nodeIdx + 1
			for idx, edge := range b.G.EdgesGoingIntoNode(nodeDest) {
				if idx == 0 {
					minOfIncomingToNodeDest = b.A[numEdgeBudget-1][edge.Tail-1] + float64(edge.Weight)
					continue
				}

				if b.A[numEdgeBudget-1][edge.Tail-1] < minOfIncomingToNodeDest {
					minOfIncomingToNodeDest = b.A[numEdgeBudget-1][edge.Tail-1] + float64(edge.Weight)
				}
			}

			if minOfIncomingToNodeDest < b.A[numEdgeBudget][nodeIdx] {
				b.A[numEdgeBudget][nodeIdx] = minOfIncomingToNodeDest
			}
		}
	}
}

func InitBellmanFord(g Graph.Graph, sourceVertex int) BellmanFord {
	b := BellmanFord{
		SourceVertex: sourceVertex,
		G:            g,
		A:            make([][]float64, g.NumVertices+1),
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
