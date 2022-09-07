package main

import "fmt"

func change(x map[int]int) { x[1], x[2] = 11, 22 }

func main() {
	x := map[int]int{0: 1, 1: 2, 2: 3, 3: 4, 4: 5}
	change(x)
	fmt.Println(x) // Output: [0:1 1:22 2:33 3:4 4:5]   The map did change, so I can conclude that maps in Go are references.
}
