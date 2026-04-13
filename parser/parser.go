package parser

import (
	"fmt"
	"os"

	"lem-in/graph"
)

func ParseFile(filename string) (*graph.Colony, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("Error: invalid data format, could not open file")
	}
	defer file.Close()
	colony := &graph.Colony{
		Rooms: make(map[string]*graph.Room),
	}
	return colony, nil
}
