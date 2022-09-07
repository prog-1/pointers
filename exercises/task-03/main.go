package main

import "fmt"

func appendf(x *[]int, y *[]int) { *x = append(*x, *y...) }

func main() {
	x := []int{1, 2, 3, 4, 5}
	y := []int{6, 7, 8, 9, 10}
	appendf(&x, &y)
	fmt.Println(x) //Output: [1 2 3 4 5 6 7 8 9 10]
}
