package main

import "fmt"

func appendToSlice(x *[]int) { *x = append(*x, 6, 7) }

func main() {
	x := []int{1, 2, 3, 4, 5}
	// p := &x
	// appendToSlice(p)
	appendToSlice(&x)
	fmt.Println(x) // Output: [1 2 3 4 5 6 7]
}
