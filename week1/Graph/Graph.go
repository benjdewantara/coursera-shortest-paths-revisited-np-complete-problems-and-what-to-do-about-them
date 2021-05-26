package Graph

import (
    "io/ioutil"
    "strconv"
    "strings"
)

type Graph struct {
    NumVertices int
    Edges       []Edge
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
            numEdges, _ := strconv.Atoi(splitStr[1])

            g.NumVertices = numVertices
            g.Edges = make([]Edge, 0, numEdges)

            continue
        }

        tail, _ := strconv.Atoi(splitStr[0])
        head, _ := strconv.Atoi(splitStr[1])
        length, _ := strconv.Atoi(splitStr[2])
        e := Edge{
            Tail:   tail,
            Head:   head,
            Weight: length,
        }

        g.Edges = append(g.Edges, e)
    }

    return g
}
