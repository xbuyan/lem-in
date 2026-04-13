package parser

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	antCountRead := false

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
		if !antCountRead {

			count, err := strconv.Atoi(line)

			if err != nil || count <= 0 {
				return nil, fmt.Errorf("Error: invalid data format, invalid number of ants")
			}
			colony.NumAnts = count
			antCountRead = true
			continue
		}

		if strings.Contains(line, "-") {
			parts := strings.Split(line, "-")

			if len(parts) != 2 {
				return nil, fmt.Errorf("ERROR: invalid data format, invalid tunnel")
			}
			room1, exists1 := colony.Rooms[parts[0]]
			room2, exists2 := colony.Rooms[parts[1]]

			if !exists1 || !exists2 {
				return nil, fmt.Errorf("Error: invalid data format, tunnel to unknown room")
			}
			room1.Connections = append(room1.Connections, room2)
			room2.Connections = append(room2.Connections, room1)
			continue
		}
		// handle room definitions
		parts := strings.Fields(line)
		if len(parts) == 3 {
			x, errX := strconv.Atoi(parts[1])
			y, errY := strconv.Atoi(parts[2])
			if errX != nil || errY != nil {
				return nil, fmt.Errorf("ERROR: invalid data format, invalid room coordinates")
			}
			room := &graph.Room{
				Name: parts[0],
				X:    x,
				Y:    y,
			}
			if nextIsStart {
				colony.Start = room
				nextIsStart = false
			}
			if nextIsEnd {
				colony.End = room
				nextIsEnd = false
			}
			colony.Rooms[room.Name] = room
			continue
		}

	}
	if colony.Start == nil {
		return nil, fmt.Errorf("Error: invalid data format, no start room found")
	}
	if colony.End == nil {
		return nil, fmt.Errorf("Error: invalid data format, no end room found")
	}
	if colony.NumAnts == 0 {
		return nil, fmt.Errorf("Error: invalid data format, invalid numer of ants")
	}
	return colony, nil
}
