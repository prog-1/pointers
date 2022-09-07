package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var value = uint32(0x11223344)
	var arr *[4]byte = (*[4]byte)(unsafe.Pointer(&value))
	for i := 0; i < 4; i++ {
		fmt.Printf("Index: %v, Value: 0x%x, Address: %v", i, arr[i], &arr[i])
		fmt.Println()
	}
}

// Output:
// Index: 0, Value: 0x44, Address: 0xc000018088
// Index: 1, Value: 0x33, Address: 0xc000018089
// Index: 2, Value: 0x22, Address: 0xc00001808a
// Index: 3, Value: 0x11, Address: 0xc00001808b

// According to output the greater the index is, the more significant is address
// (It's opposite of what I previously stated in this commit:
// https://github.com/AleksandrStukalov/pointers/commit/917a6e091217e98d42f197377cc3e5e1d015ab2f)
