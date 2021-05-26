package main

import (
    "./Graph"
    "fmt"
)

func main() {
    g := Graph.ReadTextfile("_6ff856efca965e8774eb18584754fd65_g1.txt")

    fmt.Println(g)
}
