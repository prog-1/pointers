package main

import "fmt"

func change(x [5]int)         { copy(x[1:3], []int{11, 22}) }
func changePointer(x *[5]int) { copy(x[1:3], []int{11, 22}) }

func main() {
	x := [5]int{1, 2, 3, 4, 5}
	change(x)
	fmt.Println(x) // Output: [1 2 3 4 5]
	changePointer(&x)
	fmt.Println(x) // Output: [1 11 22 4 5]
}

// Arrays are values
