package main

import "fmt"

func append2(x *[]int, y *[]int) { *x = append(*x, *y...) }

func main() {
	x := []int{1, 2, 3, 4, 5}
	y := []int{8, 9}
	append2(&x, &y)
	fmt.Println(x)
}
