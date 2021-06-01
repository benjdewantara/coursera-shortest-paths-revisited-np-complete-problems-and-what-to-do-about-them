package main

import (
	"./Graph"
	"./ShortestPath"
	"fmt"
)

func main() {
	//g := Graph.ReadTextfile("_test_g1_cycle.txt")
	//g := Graph.ReadTextfile("_test_g1.txt")
	g := Graph.ReadTextfile("_6ff856efca965e8774eb18584754fd65_g1.txt")
	bellmanFord := ShortestPath.InitBellmanFord(g, 1)
	bellmanFord.Evaluate()

	fmt.Println(bellmanFord.SourceVertex)
}
