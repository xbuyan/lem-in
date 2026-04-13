package parser

import (
	"bufio"
	"fmt"
	"os"
	"strings"

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
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// remove leading spaces
		line := scanner.Text()
		line = strings.TrimSpace(line)
		// remove empty lines
		if line == "" {
			continue
		}
		if line == "##start" {
			continue
		}
		if line == "##end" {
			continue
		}
		if strings.HasPrefix(line, "#") {
			continue
		}

	}
	return colony, nil
}
