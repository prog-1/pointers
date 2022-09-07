package main

import "fmt"

func Append(values *[]int, value int) {
	*values = append((*values), value)
}

func main() {
	values := []int{1, 2, 3, 4}
	Append(&values, 5)
	fmt.Println(values)
}
