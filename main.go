package main

import (
	"fmt"
	"os"

	"lem-in/parser"
	"lem-in/solver"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("ERROR: invalid usage, expected: go run . <filename>")
		os.Exit(1)
	}

	colony, err := parser.ParseFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	path := solver.BFS(colony)
	if path == nil {
		fmt.Println("ERROR: no path found")
		os.Exit(1)
	}

	fmt.Printf("Ants: %d\n", colony.NumAnts)
	fmt.Printf("Path found: ")
	for _, room := range path {
		fmt.Printf("%s ", room.Name)
	}
	fmt.Println()
}
