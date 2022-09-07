package main

import "fmt"

func replace(values map[string]int, key string, value int) {
	values[key] = value
}

func main() {
	values := map[string]int{"Sanja": 17, "Vay4ie": 17}
	replace(values, "Sanja", 18)
	fmt.Println(values)

	// Works, which means, that maps are pointers
}
