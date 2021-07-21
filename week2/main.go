package main

import (
	"./Graph"
	"fmt"
)

func main() {
	g := Graph.ReadTextfile("_f702b2a7b43c0d64707f7ab1b4394754_tsp.txt")
	if g.VertexCoordinate != nil {
		fmt.Println("Hell on earth")
	}
}
