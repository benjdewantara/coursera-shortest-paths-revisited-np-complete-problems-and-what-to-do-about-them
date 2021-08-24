package Graph

import "sort"

type DFSTopologicalOrderGenerator struct {
	Gr                   Graph
	CurrentSourceVertex  int
	Visited              []bool
	FinishingTime        []int
	CounterFinishingTime int
}

func (g *DFSTopologicalOrderGenerator) DFSLoopFirst() {
	for vertex := g.Gr.NumVertices; vertex >= 1; vertex-- {
		vertexName := vertex

		vertexIndx := g.Gr.getIndexOfVertex(vertexName)
		if !g.Visited[vertexIndx] {
			g.CurrentSourceVertex = vertexName
			g.DFSFirstLabeler(vertexName)
			g.FinishingTime[vertexIndx] = g.CounterFinishingTime
			g.CounterFinishingTime++
		}
	}

	for vertex := g.Gr.NumVertices; vertex >= 1; vertex-- {
		vertexName := -vertex

		vertexIndx := g.Gr.getIndexOfVertex(vertexName)
		if !g.Visited[vertexIndx] {
			g.CurrentSourceVertex = vertexName
			g.DFSFirstLabeler(vertexName)
			g.FinishingTime[vertexIndx] = g.CounterFinishingTime
			g.CounterFinishingTime++
		}
	}
}

func (g *DFSTopologicalOrderGenerator) DFSFirstLabeler(rootVertex int) {
	rootVertexIndx := g.Gr.getIndexOfVertex(rootVertex)
	if g.Visited[rootVertexIndx] {
		return
	}

	g.Visited[rootVertexIndx] = true

	if g.Gr.Adj[rootVertexIndx] == nil {
		return
	}

	for _, vertex := range g.Gr.Adj[rootVertexIndx] {
		vertexIndx := g.Gr.getIndexOfVertex(vertex)
		if !g.Visited[vertexIndx] {
			g.DFSFirstLabeler(vertex)

			g.FinishingTime[vertexIndx] = g.CounterFinishingTime
			g.CounterFinishingTime++
		}
	}
}

func (g *DFSTopologicalOrderGenerator) SortVertexByFinishingTime() {
	sortedByFinishingTime := VertexLabelSortedByFinishingTime{
		VertexLabels:  g.Gr.VertexLabels,
		FinishingTime: g.FinishingTime,
	}

	sort.Sort(&sortedByFinishingTime)

	g.Gr.VertexLabels = sortedByFinishingTime.VertexLabels
	g.FinishingTime = sortedByFinishingTime.FinishingTime
}

func InitFromGraph(g *Graph) DFSTopologicalOrderGenerator {
	o := DFSTopologicalOrderGenerator{
		Gr:                   *g,
		CurrentSourceVertex:  1,
		Visited:              make([]bool, g.NumVertices*2),
		FinishingTime:        make([]int, g.NumVertices*2),
		CounterFinishingTime: 1,
	}

	return o
}
