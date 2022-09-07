package main

import "fmt"

func change(x [5]int) { copy(x[1:3], []int{11, 22}) }

func main() {
	x := [5]int{1, 2, 3, 4, 5}
	change(x)
	fmt.Println(x) // Output: [1 2 3 4 5], not [1 11 22 4 5]   The array does not change in any way, so I can conclude that arrays in Go are values, not references.
}

// However, in some other programming languages, such as Java and C++ (from what I have found), arrays are references (so they can be changed like this).
