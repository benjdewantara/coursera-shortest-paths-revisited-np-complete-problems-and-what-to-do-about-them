package Graph

import (
	"fmt"
	"sort"
	"strings"
)

type DFSTopologicalOrderGenerator struct {
	SccLeaders           []int
	Gr                   Graph
	CurrentSourceVertex  int
	Visited              []bool
	FinishingTime        []int
	CounterFinishingTime int

	HasEncounteredSccWithNegatedVariable bool
	SccMemberMap                         map[int]bool
	SccTraversal                         strings.Builder
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

func (g *DFSTopologicalOrderGenerator) DFSLoopSecond() {
	g.Visited = make([]bool, g.Gr.NumVertices*2)
	g.HasEncounteredSccWithNegatedVariable = false

	for _, vertex := range g.Gr.VertexLabels {
		vertexIndx := g.Gr.getIndexOfVertex(vertex)
		if !g.Visited[vertexIndx] {
			g.SccLeaders = append(g.SccLeaders, vertex)
			g.SccMemberMap = make(map[int]bool)
			g.SccMemberMap[vertex] = true
			g.SccTraversal.Reset()
			g.DFSSecondTraverseScc(vertex)

			if g.HasEncounteredSccWithNegatedVariable {
				fmt.Println(fmt.Sprintf("Encountered an Scc containing both variable and negated variable"))
				fmt.Println(fmt.Sprintf("SccTraversal is as follows"))
				fmt.Println(g.SccTraversal.String())
				break
			}
		}
	}
}

func (g *DFSTopologicalOrderGenerator) DFSSecondTraverseScc(rootVertex int) {
	rootVertexIndx := g.Gr.getIndexOfVertex(rootVertex)
	if g.Visited[rootVertexIndx] {
		return
	}

	g.Visited[rootVertexIndx] = true

	_, memberMapExists := g.SccMemberMap[rootVertex]
	if !memberMapExists {
		g.SccMemberMap[rootVertex] = true
	}

	_, negatedMemberExists := g.SccMemberMap[-rootVertex]
	if negatedMemberExists {
		g.HasEncounteredSccWithNegatedVariable = true
	}

	if g.Gr.Adj[rootVertexIndx] == nil {
		return
	}

	for _, vertex := range g.Gr.Adj[rootVertexIndx] {
		vertexIndx := g.Gr.getIndexOfVertex(vertex)
		if !g.Visited[vertexIndx] {
			g.SccTraversal.WriteString(fmt.Sprintf("%d -> %d\n", rootVertex, vertex))
			g.DFSSecondTraverseScc(vertex)
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
