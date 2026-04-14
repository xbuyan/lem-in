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

func BFS(colony *graph.Colony) []*graph.Room {
	visited := map[string]bool{}
	previous := map[string]*graph.Room{}
	queue := []*graph.Room{colony.Start}
	visited[colony.Start.Name] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current == colony.End {
			return reconstructPath(previous, colony.Start, current)
		}

		for _, neighbor := range current.Connections {
			if !visited[neighbor.Name] {
				visited[neighbor.Name] = true
				previous[neighbor.Name] = current
				queue = append(queue, neighbor)
			}
		}
	}

	return nil
}
