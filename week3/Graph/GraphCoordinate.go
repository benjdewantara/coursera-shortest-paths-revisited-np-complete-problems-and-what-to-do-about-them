package Graph

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strconv"
	"strings"
)

type GraphCoordinate struct {
	Vertices         []int
	VertexCoordinate [][2]float64

	Visited            []int
	Unvisited          []int
	IndicesInUnvisited []int

	MinDist float64

	pivotVertex int
	pivotIndex  int
}

func (g *GraphCoordinate) Len() int {
	return len(g.Unvisited)
}

func (g *GraphCoordinate) Less(i int, j int) bool {
	vertexI := g.Unvisited[i]
	vertexJ := g.Unvisited[j]

	distanceI := g.SquaredDistanceBetween(g.pivotVertex, vertexI)
	distanceJ := g.SquaredDistanceBetween(g.pivotVertex, vertexJ)

	if distanceI == distanceJ {
		return vertexI < vertexJ
	}

	return distanceI < distanceJ
}

func (g *GraphCoordinate) Swap(i int, j int) {
	g.Unvisited[i], g.Unvisited[j] =
		g.Unvisited[j], g.Unvisited[i]
}

func (g *GraphCoordinate) EvaluateTsp() {
	for len(g.Unvisited) >= 2 {
		sort.Sort(g)
		g.Visited = append(g.Visited, g.Unvisited[0])
		g.Unvisited = g.Unvisited[1:]
		g.pivotVertex = g.Visited[len(g.Visited)-1]
	}

	g.Visited = append(g.Visited, g.Unvisited[0])
	g.Unvisited = g.Unvisited[1:]
	g.pivotVertex = g.Visited[len(g.Visited)-1]

	dist := 0.0
	for idx := 1; idx < len(g.Visited); idx++ {
		dist += g.DistanceBetween(g.Visited[idx], g.Visited[idx-1])
	}

	dist += g.DistanceBetween(g.Visited[len(g.Visited)-1], g.Visited[0])

	fmt.Println(fmt.Sprintf("EvaluateTsp: dist = %f", dist))
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

	g.Unvisited = append([]int{}, g.Vertices...)

	g.IndicesInUnvisited = make([]int, len(g.Unvisited))
	for indx, _ := range g.IndicesInUnvisited {
		g.IndicesInUnvisited[indx] = indx
	}

	g.Visited = append([]int{}, g.Unvisited[0])
	g.Unvisited = g.Unvisited[1:]

	g.pivotVertex = g.Visited[len(g.Visited)-1]
	g.pivotIndex = 0

	return g
}
