package main

import (
	"fmt"
	"unsafe"
)

func updateArray(array [5]int) {
	array[1] = 0
}

func updateMap(a map[int]int) {
	a[1] = 0
}

func appendSlice(slice *[]int) {
	*slice = append(*slice, 6, 7, 8, 9, 10)
}

func main() {
	// Array
	array := [5]int{1, 2, 3, 4, 5}
	updateArray(array)
	fmt.Println(array) // OUTPUT: [1 2 3 4 5]
	/*
		Function "updateArray" does not affact array in the main function so arrays are passed as are values
		As to other languages

		-	C++ http://tpcg.io/_89JPQU Arrays are references
	*/

	// Map
	a := map[int]int{1: 1, 2: 2, 3: 3, 4: 4, 5: 5}
	updateMap(a)
	fmt.Println(a) // OUTPUT: map[1:0 2:2 3:3 4:4 5:5]
	/*
		Function "updateMap" does affact map in the main function so maps are passed as reference
	*/

	// Append slice
	slice := []int{1, 2, 3, 4, 5}
	appendSlice(&slice)
	fmt.Println(slice) // OUTPUT: [1 2 3 4 5 6 7 8 9 10]

	x := uint32(0x11223344)
	c := (*[4]byte)(unsafe.Pointer(&x))
	fmt.Printf("% x\n", *c)                                         // OUTPUT: 44 33 22 11
	fmt.Printf("%x %x %x %x\n", (*c)[0], (*c)[1], (*c)[2], (*c)[3]) // OUTPUT: 44 33 22 11

	fmt.Printf("%x %v\n", x, x) // OUTPUT: 11223344 287454020
	(*c)[3] = 0x12
	fmt.Printf("%x %v\n", x, x) // OUTPUT: 11223345 304231236

	// lower addres is most significant
}
