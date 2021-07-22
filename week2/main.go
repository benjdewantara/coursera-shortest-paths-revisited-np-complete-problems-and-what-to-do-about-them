package main

import (
	"./Graph"
)

func main() {
	g := Graph.ReadTextfile("_f702b2a7b43c0d64707f7ab1b4394754_tsp.txt")
	//g := Graph.ReadTextfile("_test_tsp.txt")
	g.EvaluateTsp()
}
