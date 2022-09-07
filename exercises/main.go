package main

import (
	"fmt"
	"unsafe"
)

func appendS(s *[]int) { *s = append(*s, 123) }
func changeA(a [4]int) {
	a[1] = 2
}
func changeAP(a *[4]int) {
	a[1] = 2
}
func changeM(m map[string]int) {
	m["foo"] = 11
}

func main() {

	//1
	a := [4]int{1, 22, 3, 4}
	changeA(a)
	fmt.Println("no changes", a)
	changeAP(&a)
	fmt.Println("changes", a)
	// changes do not apply if array is called not by adress. That means that arrays are values.

	//2
	m := map[string]int{"foo": 1, "bar": 2, "boo": 3}
	changeM(m)
	fmt.Println(m)
	//its not allowed to use this↓↓↓ function with map

	// func changeMP(m *map[string]int) {
	// 	m["bar"] = 12345
	// }

	//Maps are references

	//
	//3

	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Default slice", slice)
	appendS(&slice)
	fmt.Println("Appended slice", slice)

	//4

	x := uint32(0x11223344)
	b := (*[4]byte)(unsafe.Pointer(&x))
	fmt.Printf("0x%x\n", x)
	fmt.Printf("most significant - %x\nleast significant- %x\n", b[3], b[0])
}
