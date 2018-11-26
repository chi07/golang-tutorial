package entities

type Book struct {
	Id 		uint	`gorm:"primary_key"`
	Name 	string  `json:"name"`
	Author 	string  `json:"author"`
	Year	int	`json:"year"`
}