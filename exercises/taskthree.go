package main

import "fmt"

func pointthree(n *[]int) {
	*n = append(*n, 2022)
}

func main() {
	n := []int{20, 202, 22}
	fmt.Println(n)
	pointthree(&n)
	fmt.Println(n)
}
