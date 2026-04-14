package solver

import "lem-in/graph"

func reconstructPath(previous map[string]*graph.Room, start *graph.Room, end *graph.Room) []*graph.Room {
	var path []*graph.Room
	current := end

	for current != start {
		path = append(path, current)
		current = previous[current.Name]
	}

	path = append(path, start)

	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}

	return path
}
