package main

import (
	"fmt"
	"math/rand"
	"time"
	"unsafe"
)

func randNum() int {
	max, min := 10, 1
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn((max - min) + min)
	if num < 2 {
		return num
	}
	return num - 1 + num - 2
}

func changeA(x [4]int) {
	x[1] = randNum()
}
func changeMap(x map[int]int) {
	x[1] = randNum()
}
func changeSlice(x *[]int) {
	*x = append(*x, randNum(), randNum()+2)
}
func main() {
	//Task 1
	a := [4]int{}
	changeA(a)
	fmt.Println(a) // In Go arrays are values, unlike java or javascript where it is reference to object
	//Task2
	m := map[int]int{0: 0, 1: 0, 2: 0}
	changeMap(m)
	fmt.Println(m) // Maps are references, they changing immediately, without return at the enf of function
	// Task 3
	s := []int{0, 0, 0, 0}
	changeSlice(&s)
	fmt.Println(s)
	//Task 4
	b := uint32(0x66557788)
	b2 := (*[4]byte)(unsafe.Pointer(&b))
	fmt.Printf("% x\n", &b2)
	fmt.Println(&b2[0], &b2[1], &b2[2], &b2[3])
	//The least significant is 0x88 beacuse it has the least impact on number value as the smallest
	//The most significant is 0x66
}
