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

	for psIndx, pst := range textfiles {
		g := Graph.ReadTextfile(pst, true)
		og := Graph.InitFromGraph(&g)
		og.DFSLoopFirst()

		og.SortVertexByFinishingTime()

		og.DFSLoopSecond()

		if len(og.SccLeaders) == g.NumVertices {
			fmt.Println(fmt.Sprintf("Problem %s does not have a strongly-connected component", psIndx))
		} else if len(og.SccLeaders) < g.NumVertices {
			fmt.Println(fmt.Sprintf("Problem %s may have strongly-connected components", psIndx))
		}
		//fmt.Println(len(g.Adj))
		break
	}
}
