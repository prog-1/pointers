package main

import "fmt"

func updateArray(array [5]int) {
	array[1] = 0
}

func main() {
	array := [5]int{1, 2, 3, 4, 5}
	updateArray(array)
	fmt.Println(array) // OUTPUT: [1 2 3 4 5]
	/*
		Function "updateArray" does not affact array in the main loop so arrays are values
		As to other languages

		-	C++ http://tpcg.io/_89JPQU Arrays are references
	*/
}
