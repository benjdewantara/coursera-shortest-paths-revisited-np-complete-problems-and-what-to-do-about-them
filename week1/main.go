package main

import (
	"./Graph"
	"./ShortestPath"
	"fmt"
)

func main() {
	//g := Graph.ReadTextfile("_test_g1_cycle.txt")
	//g := Graph.ReadTextfile("_test_g1.txt")
	g := Graph.ReadTextfile("_test_g1_negative_edge.txt")
	//g := Graph.ReadTextfile("_6ff856efca965e8774eb18584754fd65_g1.txt")
	johnson := ShortestPath.InitJohnson(g)

	fmt.Println(johnson.G.NumVertices)
}
