package Graph

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type GraphCoordinate struct {
	Vertices         []int
	VertexCoordinate [][2]float64
	A                map[string]float64
	MinDist          float64
}

func (g *GraphCoordinate) Length() int {
	return len(g.VertexCoordinate)
}

func (g *GraphCoordinate) EvaluateTsp() {
	var finalSubset []int

	for m := 2; m <= g.Length(); m++ {
		subsetGenerator := CombinationExceptFirstElem{
			Elements: g.Vertices,
			M:        m,
		}

		for _, subset := range subsetGenerator.GetCombinations() {
			finalSubset = append([]int{}, subset...)
			for jIdxInSubset := 1; jIdxInSubset < len(subset); jIdxInSubset++ {
				vertexJInSubset := subset[jIdxInSubset]
				g.putDistance(subset, vertexJInSubset)
			}
		}
	}

	minDist := -1.0
	for i := 1; i < len(g.Vertices); i++ {
		vertexSecondToLast := g.Vertices[i]
		key := SubsetWVertexDestToString(finalSubset, vertexSecondToLast)

		d := g.A[key] +
			g.DistanceBetween(vertexSecondToLast, 1)
		if minDist == -1 || d < minDist {
			minDist = d
		}
	}

	g.MinDist = minDist

	fmt.Println(fmt.Sprintf("EvaluateTsp: finalSubset = %s", finalSubset))
	fmt.Println(fmt.Sprintf("EvaluateTsp: MinDist = %f", g.MinDist))
	fmt.Println("EvaluateTsp: End of function")
}

func (g *GraphCoordinate) DistanceBetween(vertexA int, vertexB int) float64 {
	xA, yA := g.VertexCoordinate[vertexA-1][0], g.VertexCoordinate[vertexA-1][1]
	xB, yB := g.VertexCoordinate[vertexB-1][0], g.VertexCoordinate[vertexB-1][1]
	deltaX := xA - xB
	deltaY := yA - yB
	return math.Sqrt((deltaX * deltaX) + (deltaY * deltaY))
}

func (g *GraphCoordinate) putDistance(subset []int, vertexJInSubset int) float64 {
	key := SubsetWVertexDestToString(subset, vertexJInSubset)
	dist, exists := g.A[key]
	if exists {
		return dist
	}

	if len(subset) == 1 {
		dist = math.Inf(1)
		if subset[0] == 1 && vertexJInSubset == 1 {
			dist = 0
		}
		g.A[key] = dist
		return g.A[key]
	} else if len(subset) == 2 {
		dist = g.DistanceBetween(subset[0], subset[1])
		g.A[key] = dist
		return g.A[key]
	}

	minDist := -1.0
	for kIdxInSubset := 1; kIdxInSubset < len(subset); kIdxInSubset++ {
		vertexKInSubset := subset[kIdxInSubset]
		if vertexKInSubset == vertexJInSubset {
			continue
		}

		vertexDestInSubsetIdx := -1
		for destIdxInS := 0; destIdxInS < len(subset); destIdxInS++ {
			if subset[destIdxInS] == vertexJInSubset {
				vertexDestInSubsetIdx = destIdxInS
				break
			}
		}

		subsetWoVertexDest := append([]int{}, subset[0:vertexDestInSubsetIdx]...)
		subsetWoVertexDest = append(subsetWoVertexDest, subset[vertexDestInSubsetIdx+1:]...)

		d := g.putDistance(subsetWoVertexDest, vertexKInSubset) +
			g.DistanceBetween(vertexKInSubset, vertexJInSubset)
		if minDist == -1 || d < minDist {
			minDist = d
		}
	}

	g.A[key] = minDist
	return g.A[key]
}

func SubsetWVertexDestToString(subset []int, vertexDest int) string {
	return fmt.Sprintf("%s>%d", IndicesToString(subset), vertexDest)
}

func IndicesToString(indices []int) string {
	s := fmt.Sprintf("%d", indices[0])
	for i := 1; i < len(indices); i++ {
		s = fmt.Sprintf("%s,%d", s, indices[i])
	}

	return s
}

func ReadTextfile(filepath string) GraphCoordinate {
	g := GraphCoordinate{}
	g.A = make(map[string]float64)

	contentBytes, _ := ioutil.ReadFile(filepath)
	for lineIndx, intStr := range strings.Split(string(contentBytes), "\n") {
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

		x, _ := strconv.ParseFloat(splitStr[0], 64)
		y, _ := strconv.ParseFloat(splitStr[1], 64)

		vertexIndx := lineIndx - 1
		//g.VertexCoordinate[vertexIndx] = make([]int, 2)
		g.VertexCoordinate[vertexIndx][0] = x
		g.VertexCoordinate[vertexIndx][1] = y
	}

	return g
}
