package main

import (
	"github.com/benjdewantara/coursera-shortest-paths-revisited-np-complete-problems-and-what-to-do-about-them/week3/Graph"
)

func main() {
	g := Graph.ReadTextfile("_ae5a820392a02042f87e3b437876cf19_nn.txt")
	g.EvaluateTsp()
}
