package club

type Club struct { // exported struct
	Name  	string //exported field
	Country string //exported field
	couch 	string //unexported field
}