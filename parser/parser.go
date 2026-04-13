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
	nextIsStart := false
	nextIsEnd := false
	_ = nextIsStart
	_ = nextIsEnd
	for scanner.Scan() {
		// remove leading spaces
		line := scanner.Text()
		line = strings.TrimSpace(line)
		// remove empty lines
		if line == "" {
			continue
		}
		if line == "##start" {
			nextIsStart = true
			continue
		}
		if line == "##end" {
			nextIsEnd = true
			continue
		}
		if strings.HasPrefix(line, "#") {
			continue
		}

	}
	return colony, nil
}
