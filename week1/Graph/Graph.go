package Graph

import (
	"io/ioutil"
	"strconv"
	"strings"
)

type Graph struct {
	NumVertices int
	Adj         [][][2]int
}

func ReadTextfile(filepath string) Graph {
	g := Graph{}

	contentBytes, _ := ioutil.ReadFile(filepath)
	for lineIndx, intStr := range strings.Split(string(contentBytes), "\n") {
		if intStr == "" {
			continue
		}

		splitStr := strings.Split(intStr, " ")

		if lineIndx == 0 {
			numVertices, _ := strconv.Atoi(splitStr[0])
			//numEdges, _ := strconv.Atoi(splitStr[1])

			g.NumVertices = numVertices
			g.Adj = make([][][2]int, numVertices)

			continue
		}

		tail, _ := strconv.Atoi(splitStr[0])
		head, _ := strconv.Atoi(splitStr[1])
		length, _ := strconv.Atoi(splitStr[2])

		if g.Adj[tail-1] == nil {
			g.Adj[tail-1] = make([][2]int, 0)
		}

		g.Adj[tail-1] = append(g.Adj[tail-1], [2]int{head, length})
	}

	return g
}
