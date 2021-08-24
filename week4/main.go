package main

import (
	"fmt"
	"github.com/benjdewantara/coursera-shortest-paths-revisited-np-complete-problems-and-what-to-do-about-them/week4/Graph"
)

func main() {
	//g := Graph.ReadTextfile("_02c1945398be467219866ee1c3294d2d_sample2sat0.txt", false)
	//g := Graph.ReadTextfile("_02c1945398be467219866ee1c3294d2d_sample2sat1.txt", false)

	textfiles := []string{
		"_02c1945398be467219866ee1c3294d2d_2sat1.txt",
		"_02c1945398be467219866ee1c3294d2d_2sat2.txt",
		"_02c1945398be467219866ee1c3294d2d_2sat3.txt",
		"_02c1945398be467219866ee1c3294d2d_2sat4.txt",
		"_02c1945398be467219866ee1c3294d2d_2sat5.txt",
		"_02c1945398be467219866ee1c3294d2d_2sat6.txt"}

	//textfiles = []string{
	//	"_02c1945398be467219866ee1c3294d2d_sample2sat2.txt",
	//	"_02c1945398be467219866ee1c3294d2d_sample2sat3.txt"}

	bitmaskString := ""
	for psIndx, pst := range textfiles {
		fmt.Println(fmt.Sprintf("Evaluating problem %d textfile %s", psIndx+1, pst))

		gReversed := Graph.ReadTextfile(pst, true)
		og := Graph.InitFromGraph(&gReversed)
		og.DFSLoopFirst()

		og.SortVertexByFinishingTime()

		g := Graph.ReadTextfile(pst, false)
		g.VertexLabels = gReversed.VertexLabels
		og.Gr = g

		og.DFSLoopSecond()

		if og.HasEncounteredSccWithNegatedVariable {
			bitmaskString += "0"
		} else {
			bitmaskString += "1"
		}
	}

	fmt.Println()
	fmt.Println(bitmaskString)
}
