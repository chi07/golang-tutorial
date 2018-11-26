package player

type Player struct {
	Name, Position 	string //exported field
	Number, Age	    uint //exported field
	stadium string //unexported field
}