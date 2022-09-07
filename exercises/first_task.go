package main

import "fmt"

//arrays are values
func change(x [5]int)   { copy(x[1:3], []int{11, 22}) }
func change2(x *[5]int) { copy(x[1:3], []int{11, 22}) }
func main() {
	x := [...]int{1, 2, 3, 4, 5}
	x2 := [...]int{1, 2, 3, 4, 5}
	change(x)
	change2(&x2)
	fmt.Println(x)  // Output: [1 2 3 4 5]
	fmt.Println(x2) // Output: [1 11 22 4 5]
}
