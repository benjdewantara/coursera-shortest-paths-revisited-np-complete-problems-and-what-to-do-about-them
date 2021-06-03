package main

import (
	"./Graph"
	"./ShortestPath"
)

func main() {
	//g := Graph.ReadTextfile("_test_g1_cycle.txt")
	//g := Graph.ReadTextfile("_test_g1.txt")
	//g := Graph.ReadTextfile("_test_g1_negative_edge.txt")

	filenames := []string{
		"_6ff856efca965e8774eb18584754fd65_g1.txt",
		"_6ff856efca965e8774eb18584754fd65_g2.txt",
		"_6ff856efca965e8774eb18584754fd65_g3.txt",
	}

	for _, filename := range filenames {
		g := Graph.ReadTextfile(filename)
		johnson := ShortestPath.InitJohnson(g)
		johnson.PrintShortestPairIfAny()
	}
}
