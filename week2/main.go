package main

import (
	"github.com/benjdewantara/coursera-shortest-paths-revisited-np-complete-problems-and-what-to-do-about-them/week2/Graph"
)

func main() {
	g := Graph.ReadTextfile("_f702b2a7b43c0d64707f7ab1b4394754_tsp.txt")
	//g := Graph.ReadTextfile("_test_tsp.txt")
	g.EvaluateTsp()
}
