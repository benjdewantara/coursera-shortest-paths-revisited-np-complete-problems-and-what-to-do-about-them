package main

import (
	"fmt"
	"github.com/benjdewantara/coursera-shortest-paths-revisited-np-complete-problems-and-what-to-do-about-them/week4/Graph"
)

func main() {
	//g := Graph.ReadTextfile("_02c1945398be467219866ee1c3294d2d_sample2sat0.txt", false)
	//g := Graph.ReadTextfile("_02c1945398be467219866ee1c3294d2d_sample2sat1.txt", false)

	g := Graph.ReadTextfile("_02c1945398be467219866ee1c3294d2d_2sat1.txt", true)
	og := Graph.InitFromGraph(&g)
	og.DFSLoopFirst()

	og.SortVertexByFinishingTime()

	fmt.Println(len(g.Adj))
}
