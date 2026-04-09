package graph

// Room represents a single room in the ant colony

type Room struct {
	Name string

	X int

	Y int

	AntCount int

	Connections []*Room
}

// Colony represents the entire ant farm

type Colony struct {
	Rooms map[string]*Room

	Start *Room

	End *Room

	NumAnts int
}

type Ant struct {
	ID          int
	CurrentRoom []*Room
	Path        []*Room
}
