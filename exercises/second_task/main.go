package main

import "fmt"

//maps are references
func mapchange(m map[string]int) { m["Helena"] = 11 }
func main() {
	age := map[string]int{
		"Anna":   10,
		"Helena": 16,
	}
	age2 := map[string]int{
		"Anna":   10,
		"Helena": 16,
	}
	fmt.Println(age)
	mapchange(age2)
	fmt.Println(age2)
}
