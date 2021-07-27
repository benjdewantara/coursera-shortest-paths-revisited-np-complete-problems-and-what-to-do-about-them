package Graph

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type GraphCoordinate struct {
	Vertices         []int
	VertexCoordinate [][2]float64
	A                map[string]float64
	ABitIndexed      []float64
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
		//key := SubsetWVertexDestToString(finalSubset, vertexSecondToLast)
		//val, _ := GetValueFromTextfile(key)

		bitIndex := BitIndex(finalSubset)
		d := g.ABitIndexed[bitIndex] +
			g.DistanceBetween(vertexSecondToLast, 1)
		if minDist == -1 || d < minDist {
			minDist = d
		}
	}

	g.MinDist = minDist

	fmt.Println(finalSubset)
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
	//key := SubsetWVertexDestToString(subset, vertexJInSubset)
	//dist, exists := GetValueFromTextfile(key)
	//dist, exists := g.A[key]
	//if exists {
	//	return dist
	//}

	bitIndex := BitIndex(subset)
	dist := g.ABitIndexed[bitIndex]
	if dist != 0 {
		return g.ABitIndexed[bitIndex]
	}

	if len(subset) == 1 {
		dist = math.Inf(1)
		if subset[0] == 1 && vertexJInSubset == 1 {
			dist = 0
		}
		g.ABitIndexed[bitIndex] = dist

		//WriteValueToTextfile(key, dist)
		//g.A[key] = dist

		return dist
	} else if len(subset) == 2 {
		dist = g.DistanceBetween(subset[0], subset[1])

		g.ABitIndexed[bitIndex] = dist

		//WriteValueToTextfile(key, dist)
		//g.A[key] = dist

		return dist
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

	g.ABitIndexed[bitIndex] = minDist

	//WriteValueToTextfile(key, minDist)
	//g.A[key] = minDist

	return minDist
}

func SubsetWVertexDestToString(subset []int, vertexDest int) string {
	return fmt.Sprintf("%s>%d", IndicesToString(subset), vertexDest)
}

func BitIndex(indices []int) int {
	summed := 0
	for _, vertex := range indices {
		summed += 1 << (vertex - 1)
	}
	return summed
}

func IndicesToString(indices []int) string {
	s := fmt.Sprintf("%d", indices[0])
	for i := 1; i < len(indices); i++ {
		s = fmt.Sprintf("%s,%d", s, indices[i])
	}

	return s
}

func GetValueFromTextfile(key string) (float64, bool) {
	return -1.0, false

	directoryFullpath := "/media/e/Benjamin Antara/Documents/Online Courses/Coursera/coursera-textfile-data2"
	directoryFullpath = strings.TrimRight(directoryFullpath, "/")

	filenamePostfix := strings.Split(key, ">")[0]
	filename := fmt.Sprintf("%s/values%s.txt", directoryFullpath, filenamePostfix)

	file, err := os.Open(filename)
	if err != nil {
		file, err = os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0755)
	}
	defer file.Close()

	num := 0.0
	exists := false
	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example

	for scanner.Scan() {
		s := scanner.Text()

		if s == "" {
			continue
		}

		splitStr := strings.Split(s, " ")
		if splitStr[0] == key {
			num, _ = strconv.ParseFloat(splitStr[1], 64)
			exists = true
			break
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return num, exists
}

func WriteValueToTextfile(key string, value float64) {
	return

	directoryFullpath := "/media/e/Benjamin Antara/Documents/Online Courses/Coursera/coursera-textfile-data2"
	directoryFullpath = strings.TrimRight(directoryFullpath, "/")

	filenamePostfix := strings.Split(key, ">")[0]
	filename := fmt.Sprintf("%s/values%s.txt", directoryFullpath, filenamePostfix)

	file, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND, 0755)
	if err != nil {
		file, err = os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0755)
	}
	defer file.Close()

	file.WriteString(fmt.Sprintf("%s %f\n", key, value))
}

func ReadTextfile(filepath string) GraphCoordinate {
	g := GraphCoordinate{}
	//g.A = make(map[string]float64)

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

			powerSetSize := math.Pow(2, float64(numVertices))
			g.ABitIndexed = make([]float64, int(powerSetSize))

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
