package main

import (
    "./Graph"
    "./ShortestPath"
    "fmt"
)

func main() {
    g := Graph.ReadTextfile("_6ff856efca965e8774eb18584754fd65_g1.txt")
    bellmanFord := ShortestPath.InitBellmanFord(g, 1)

    fmt.Println(bellmanFord.SourceVertex)
}
