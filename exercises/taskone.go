package main

import (
	"fmt"
)

func pointone(n *[5]int) {
	copy(n[1:4], []int{00, 00, 00})
}
func main() {
	n := []int{99, 88, 77, 66, 55}
	pointone(&n)
	fmt.Println(n)
	// arrays are values
}
