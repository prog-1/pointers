package main

import (
	"fmt"
	"unsafe"
)

// func change(a *[5]int) { copy(a[1:3], []int{11, 22}) }
// func main() {
// 	a := [5]int{1, 2, 3, 4, 5}
// 	change(&a)
// 	fmt.Println(a) // Output: [1 11 22 4 5]
// } //Arrays are values

// func change(m map[int]int) { m[1] = 5 }
// func main() {
// 	m := make(map[int]int)
// 	m[1] = 1
// 	m[2] = 2
// 	m[3] = 3
// 	change(m)
// 	fmt.Println(m) // Output: map[1:5 2:2 3:3]
// } //Maps are references

// func change(x *[]int) { *x = append(*x, 6) }
// func main() {
// 	x := []int{1, 2, 3, 4, 5}
// 	change(&x)
// 	fmt.Println(x) // Output: [1 2 3 4 5 6]
// }

func main() {
	x := uint32(0x11223344)
	a := (*[4]byte)(unsafe.Pointer(&x))
	fmt.Printf("%x %v\n", a[0], &a[0]) // Output: 44 0xc000102058
	fmt.Printf("%x %v\n", a[3], &a[3]) // Output: 11 0xc00010205b
} //lower address is the least significant