package Graph

import (
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type GraphCoordinate struct {
	Vertices         []int
	VertexCoordinate [][2]float64
	MinDist          float64
}

func (g *GraphCoordinate) Length() int {
	return len(g.VertexCoordinate)
}

func (g *GraphCoordinate) EvaluateTsp() {

}

func (g *GraphCoordinate) SquaredDistanceBetween(vertexA int, vertexB int) float64 {
	xA, yA := g.VertexCoordinate[vertexA-1][0], g.VertexCoordinate[vertexA-1][1]
	xB, yB := g.VertexCoordinate[vertexB-1][0], g.VertexCoordinate[vertexB-1][1]
	deltaX := xA - xB
	deltaY := yA - yB
	return (deltaX * deltaX) + (deltaY * deltaY)
}

func (g *GraphCoordinate) DistanceBetween(vertexA int, vertexB int) float64 {
	return math.Sqrt(g.SquaredDistanceBetween(vertexA, vertexB))
}

func ReadTextfile(filepath string) GraphCoordinate {
	g := GraphCoordinate{}
	//g.A = make(map[string]float64)

	contentBytes, _ := ioutil.ReadFile(filepath)
	for lineIndx, intStr := range strings.Split(string(contentBytes), "\n") {
		intStr = strings.TrimRight(intStr, "\r\n")

		if intStr == "" {
			continue
		}

		splitStr := strings.Split(intStr, " ")

		if lineIndx == 0 {
			numVertices, _ := strconv.Atoi(splitStr[0])
			g.VertexCoordinate = make([][2]float64, numVertices)

			g.Vertices = make([]int, numVertices)
			for vidx := 0; vidx < len(g.Vertices); vidx++ {
				g.Vertices[vidx] = vidx + 1
			}

			continue
		}

		x, _ := strconv.ParseFloat(splitStr[1], 64)
		y, _ := strconv.ParseFloat(splitStr[2], 64)

		vertexIndx := lineIndx - 1
		//g.VertexCoordinate[vertexIndx] = make([]int, 2)
		g.VertexCoordinate[vertexIndx][0] = x
		g.VertexCoordinate[vertexIndx][1] = y
	}

	return g
}
