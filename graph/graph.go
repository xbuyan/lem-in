package graph

type Room struct {
	Name string

	X int

	Y int

	AntCount int

	Connections []*Room
}

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
	Step        int
}
