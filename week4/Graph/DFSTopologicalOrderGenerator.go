package Graph

type DFSTopologicalOrderGenerator struct {
	Gr            Graph
	Visited       []bool
	FinishingTime int
}

func (g *DFSTopologicalOrderGenerator) DFSLoop() {
	for vertex := 1; vertex < g.Gr.NumVertices; vertex++ {

	}
}

func InitFromGraph(g *Graph) DFSTopologicalOrderGenerator {
	o := DFSTopologicalOrderGenerator{
		Gr:            *g,
		Visited:       make([]bool, g.NumVertices*2),
		FinishingTime: 1,
	}

	return o
}
