package main

import (
	"fmt"
	"unsafe"
)

func main() {
	dWVar := uint32(0x11223344)               //making the double word variable
	Arr := (*[4]byte)(unsafe.Pointer(&dWVar)) // apllying [4]byte array on top of dwVar

	for i := 0; i < len(Arr); i++ {
		fmt.Println("Index: ", i, " Value: ", Arr[i], "Address: ", &Arr[i])
	}

	//Conclusion: Addresses of the Values goes from the most significant to the less significant. Adress from the Value of the Index 1 will be the Most significant, whereas Adress from the Last Index will be Less significant.
}
