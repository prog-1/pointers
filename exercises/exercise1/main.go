package main

import "fmt"

func Replace(array [4]int, replaceIndex int, replaceValue int) {
	array[replaceIndex] = replaceValue
}

func main() {
	var arr = [4]int{1, 2, 3, 4}
	Replace(arr, 1, 0)
	fmt.Println(arr)

	// Nothing changes, which means, that, unlike c-style languages, array is not a pointer, but value
}
