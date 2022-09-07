package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var value = uint32(0x11223344)
	var ptr *[4]byte = (*[4]byte)(unsafe.Pointer(&value))
	for i := 0; i < 4; i++ {
		fmt.Println("Index:", i, "\t", "Address:", &ptr[i])
	}

	// Output:
	// Index: 0		Address: 0xc0000a6058
	// Index: 1		Address: 0xc0000a6059
	// Index: 2		Address: 0xc0000a605a
	// Index: 3		Address: 0xc0000a605b

	// According to output, the greater the byte index is, the less significant is its address and vice versa
}
