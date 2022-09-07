package main

import (
	"fmt"
	"unsafe"
)

func main() {
	x := uint32(0x11223344)
	a := (*[4]byte)(unsafe.Pointer(&x))
	fmt.Printf("% x\n", *a)
	fmt.Println(&a[0], &a[1], &a[2], &a[3])
	//the least significant has a lower address
}
