package main

import (
	"fmt"
	"os"

	"lem-in/parser"
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
	fmt.Printf("Ants: %d\n", colony.NumAnts)
	fmt.Printf("Rooms: %d\n", len(colony.Rooms))
	fmt.Printf("Start: %s\n", colony.Start.Name)
	fmt.Printf("End: %s\n", colony.End.Name)
}
