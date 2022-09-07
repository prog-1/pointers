package main

import "fmt"

func pointtwo(s map[string]int) {
	s["Maya"] = 3012
}

func main() {
	Doomsday := map[string]int{
		"Maya":    2012,
		"Ice_age": 9999,
	}
	fmt.Println(Doomsday)
	pointtwo(Doomsday)
	fmt.Println(Doomsday)
	//Maps are references
}
