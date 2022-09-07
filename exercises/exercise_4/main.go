package main

import (
	"fmt"
	"unsafe"
)

func main() {
	x := uint32(0x11223344)
	a := (*[4]byte)(unsafe.Pointer(&x))
	fmt.Printf("% x\n", *a)                                         // Output: 44 33 22 11
	fmt.Printf("%x %x %x %x\n", (*a)[0], (*a)[1], (*a)[2], (*a)[3]) // Output: 44 33 22 11
	fmt.Println((*a)[0], (*a)[1], (*a)[2], (*a)[3])                 // Output: 68 51 34 17 (in decimal)
	fmt.Println(&((*a)[0]), &((*a)[1]), &((*a)[2]), &((*a)[3]))     // byte addresses
}

// Byte adresses are always in ascending order (or aplphabetical order if there there are letters at the end), so the most significant byte (in out case it's 0x11) has the highest adress.
// That's why the least significant byte (in our case it's 0x44) has a lower adress.
