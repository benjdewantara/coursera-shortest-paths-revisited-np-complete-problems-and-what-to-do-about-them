package ShortestPath

import (
	"../Graph"
	"container/heap"
	"math"
)

type Dijkstra struct {
	G                     Graph.Graph
	RootNode              int
	ShortestLengthToNodes []float64
	AdjacentNodes         []DistanceFromRoot
	SuperNodes            map[int]bool
}

func (d *Dijkstra) Evaluate() {
	for len(d.AdjacentNodes) > 0 {
		distFromRoot := heap.Pop(d).(DistanceFromRoot)
		_, exists := d.SuperNodes[distFromRoot.NodeDestination]
		if exists {
			continue
		}

		heads := d.G.Adj[distFromRoot.NodeDestination-1]
		for _, h := range heads {
			nodeDest := h[0]
			weight := float64(h[1]) + distFromRoot.Distance

			newDist := DistanceFromRoot{nodeDest, weight}
			heap.Push(d, newDist)
		}
	}
}

func InitDijkstraWithRoot(g Graph.Graph, root int) Dijkstra {
	d := Dijkstra{
		G:                     g,
		RootNode:              root,
		ShortestLengthToNodes: make([]float64, g.NumVertices),
		AdjacentNodes:         make([]DistanceFromRoot, 0),
		SuperNodes:            make(map[int]bool),
	}

	for nodeIdx := 0; nodeIdx < len(d.ShortestLengthToNodes); nodeIdx++ {
		d.ShortestLengthToNodes[nodeIdx] = math.Inf(1)
	}
	d.ShortestLengthToNodes[root-1] = float64(0)
	d.AdjacentNodes = append(d.AdjacentNodes, DistanceFromRoot{d.RootNode, 0})
	//d.AdjacentNodes[root-1].Distance = 0
	//d.AdjacentNodes = append(d.AdjacentNodes[0:root-1], d.AdjacentNodes[root:]...)
	//d.SuperNodes[root] = true

	heap.Init(&d)
	return d
}

func (d *Dijkstra) Len() int {
	return len(d.AdjacentNodes)
}

func (d *Dijkstra) Less(i, j int) bool {
	return d.AdjacentNodes[i].Distance < d.AdjacentNodes[i].Distance
}

func (d *Dijkstra) Swap(i, j int) {
	d.AdjacentNodes[i], d.AdjacentNodes[j] = d.AdjacentNodes[j], d.AdjacentNodes[i]
}

func (d *Dijkstra) Push(x interface{}) {
	d.AdjacentNodes = append(d.AdjacentNodes, x.(DistanceFromRoot))
}

func (d *Dijkstra) Pop() interface{} {
	lastElem := d.AdjacentNodes[len(d.AdjacentNodes)-1]
	d.AdjacentNodes = d.AdjacentNodes[0 : len(d.AdjacentNodes)-1]
	return lastElem
}

type DistanceFromRoot struct {
	NodeDestination int
	Distance        float64
}
