package main

import (
	"fmt"
	"golang-tutorial/struct/club"
)

func main()  {
	club  := club.Club{}
	club.Name = "ManUtd"
	club.Country = "England"
	club.couch = "Jose Mourinho"
	fmt.Println(club)
}