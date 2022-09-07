package main

import "fmt"

func change(x [5]int) { copy(x[1:3], []int{11, 22}) }

func change2(y *[5]int) { copy(y[1:3], []int{11, 22}) }

func main() {
	x := [...]int{1, 2, 3, 4, 5}
	y := [...]int{1, 2, 3, 4, 5}
	change(x)
	change2(&y)
	fmt.Println(x) // Output: [1 2 3 4 5]
	fmt.Println(y) // Output: [1 11 22 4 5]
}

//arrays are values in golang, but in language like C, C++, arrays are references
//https://www.codingninjas.com/blog/2021/08/31/passing-arrays-to-functions-in-c-c/#:~:text=Passing%20arrays%20to%20functions%20in%20C%2FC%2B%2B%20are%20passed%20by,back%20to%20the%20original%20array.
