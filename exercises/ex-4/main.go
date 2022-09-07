package main

import (
	"fmt"
	"unsafe"
)

func main() {
	x := uint32(0x11223344)
	a := (*[4]byte)(unsafe.Pointer(&x))
	fmt.Printf("%x %v\n", a[0], &a[0]) // Output : 44 0xc000128058
	fmt.Printf("%x %v\n", a[3], &a[3]) // Output : 11 0xc00012805b
}

// Lower address is the least significant
