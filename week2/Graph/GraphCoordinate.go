package Graph

import (
	"io/ioutil"
	"strconv"
	"strings"
)

type GraphCoordinate struct {
	VertexCoordinate [][2]float64
}

func ReadTextfile(filepath string) GraphCoordinate {
	g := GraphCoordinate{}

	contentBytes, _ := ioutil.ReadFile(filepath)
	for lineIndx, intStr := range strings.Split(string(contentBytes), "\n") {
		if intStr == "" {
			continue
		}

		splitStr := strings.Split(intStr, " ")

		if lineIndx == 0 {
			numVertices, _ := strconv.Atoi(splitStr[0])
			g.VertexCoordinate = make([][2]float64, numVertices)
			continue
		}

		x, _ := strconv.ParseFloat(splitStr[0], 64)
		y, _ := strconv.ParseFloat(splitStr[1], 64)

		vertexIndx := lineIndx - 1
		//g.VertexCoordinate[vertexIndx] = make([]int, 2)
		g.VertexCoordinate[vertexIndx][0] = x
		g.VertexCoordinate[vertexIndx][1] = y
	}

	return g
}
