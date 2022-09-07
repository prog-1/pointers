package main

import "fmt"

func change(m map[int]int) { m[0] = 0 }

func main() {
	m := map[int]int{
		0: 1,
		1: 1,
	}
	change(m)
	fmt.Println(m) // Output: map[0:0 1:1]
}

// Maps are references
