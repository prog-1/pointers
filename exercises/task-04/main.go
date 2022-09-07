package main

import (
	"fmt"
	"unsafe"
)

func main() {
	x := uint32(0x11223344)
	b := (*[4]byte)(unsafe.Pointer(&x))
	fmt.Printf("% x\n", *b)                                 // Output: 44 33 22 11
	fmt.Printf("%x %x %x %x\n", &b[0], &b[1], &b[2], &b[3]) // Output: c000018098 c000018099 c00001809a c00001809b

}

// lower address is the most significant
