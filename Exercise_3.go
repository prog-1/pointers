package main

import "fmt"

func Append(slice []int, index int, value int) {
	slice[index] = value
}

func main3() {
	someSlice := []int{3, 5, 2, 4}
	Append(someSlice, 1, 1)
	fmt.Println(someSlice)
}

//Conclusion: The slice also changes, trying to change the values of it, because, how is written in prog-1/pointers, slices are references too.
