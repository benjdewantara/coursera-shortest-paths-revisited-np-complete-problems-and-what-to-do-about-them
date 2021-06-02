package ShortestPath

import (
	"../Graph"
)

type Johnson struct {
	G           Graph.Graph
	GPrime      Graph.Graph
	NodeExtra   int
	BellmanFord BellmanFord
}

func InitJohnson(g Graph.Graph) Johnson {
	gPrime := g
	gPrime.Adj = append(gPrime.Adj, make([][2]int, g.NumVertices))
	gPrime.NumVertices++

	for nodeIdx := 0; nodeIdx < g.NumVertices; nodeIdx++ {
		node := nodeIdx + 1
		gPrime.Adj[len(gPrime.Adj)-1][nodeIdx] = [2]int{node, 0}
	}

	var j = Johnson{
		G:         g,
		GPrime:    gPrime,
		NodeExtra: g.NumVertices + 1,
	}

	j.BellmanFord = InitBellmanFord(gPrime, j.NodeExtra)
	j.BellmanFord.Evaluate()

	if !j.BellmanFord.NegativeCycleExists {
		for tailIdx := 0; tailIdx < j.G.NumVertices; tailIdx++ {
			for headIdx := 0; headIdx < len(j.GPrime.Adj[tailIdx]); headIdx++ {
				if j.GPrime.Adj[tailIdx] == nil {
					continue
				}

				e := Graph.Edge{
					Tail:   tailIdx + 1,
					Head:   j.GPrime.Adj[tailIdx][headIdx][0],
					Weight: j.GPrime.Adj[tailIdx][headIdx][1],
				}

				j.GPrime.Adj[tailIdx][headIdx][1] =
					e.Weight +
						int(j.BellmanFord.A[gPrime.NumVertices][tailIdx]-j.BellmanFord.A[gPrime.NumVertices][e.Head-1])
			}
		}
	}

	// get rid of the extra node
	j.GPrime.Adj = j.GPrime.Adj[0 : len(j.GPrime.Adj)-1]
	j.GPrime.NumVertices--

	if !j.BellmanFord.NegativeCycleExists {
		d := InitDijkstraWithRoot(j.GPrime, 1)
		d.Evaluate()
	}

	return j
}
