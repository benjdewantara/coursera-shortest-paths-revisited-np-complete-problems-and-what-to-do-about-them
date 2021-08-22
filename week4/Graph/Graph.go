package Graph

import (
	"io/ioutil"
	"strconv"
	"strings"
)

type Graph struct {
	NumVertices int
	Adj         [][]int
}

func ReadTextfile(filepath string) Graph {
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
			g.Adj = make([][]int, g.NumVertices*2)
			continue
		}

		vertexFrom, _ := strconv.Atoi(splitStr[0])
		vertexTo, _ := strconv.Atoi(splitStr[1])

		vertexFromIdx := vertexFrom - 1
		if vertexFromIdx < 0 {
			vertexFromIdx = (vertexFrom + g.NumVertices) - 1
		}

		vertexToIdx := vertexTo - 1
		if vertexToIdx < 0 {
			vertexToIdx = (vertexTo + g.NumVertices) - 1
		}

		if g.Adj[vertexFromIdx] == nil {
			g.Adj[vertexFromIdx] = make([]int, 1)
		}

		g.Adj[vertexFromIdx] = append(g.Adj[vertexFromIdx], vertexTo)
	}

	return g
}
