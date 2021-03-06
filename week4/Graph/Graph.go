package Graph

import (
	"io/ioutil"
	"strconv"
	"strings"
)

type Graph struct {
	VertexLabels []int
	NumVertices  int
	Adj          [][]int
}

func (g *Graph) getIndexOfVertex(vertex int) int {
	if vertex < 0 {
		return (-vertex + g.NumVertices) - 1
	}

	return vertex - 1
}

func (g *Graph) getVertexFromIndex(indx int) int {
	if indx < g.NumVertices {
		return indx + 1
	}

	return -(indx - g.NumVertices + 1)
}

func ReadTextfile(filepath string, isReversed bool) Graph {
	g := Graph{}

	contentBytes, _ := ioutil.ReadFile(filepath)
	for lineIndx, intStr := range strings.Split(string(contentBytes), "\n") {
		intStr = strings.TrimRight(intStr, "\r\n")
		splitStr := strings.Split(intStr, " ")

		if intStr == "" {
			continue
		}

		if lineIndx == 0 {
			g.NumVertices, _ = strconv.Atoi(splitStr[0])
			g.VertexLabels = make([]int, g.NumVertices*2)
			g.Adj = make([][]int, g.NumVertices*2)
			continue
		}

		vertexFrom, _ := strconv.Atoi(splitStr[0])
		vertexTo, _ := strconv.Atoi(splitStr[1])

		if isReversed {
			vertexFrom, vertexTo = vertexTo, vertexFrom
		}

		vertexFrom = -vertexFrom
		vertexFromIdx := g.getIndexOfVertex(vertexFrom)

		g.Adj[vertexFromIdx] = append(g.Adj[vertexFromIdx], vertexTo)
	}

	for i := 0; i < g.NumVertices; i++ {
		vertex := i + 1
		g.VertexLabels[i] = vertex

		vertex = -vertex
		g.VertexLabels[g.NumVertices+i] = vertex
	}

	return g
}
