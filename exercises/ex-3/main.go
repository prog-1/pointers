package main

import "fmt"

func appendPointer(s *[]int, e []int) { *s = append(*s, e...) }

func main() {
	s := []int{1, 2, 3, 4, 5}
	appendPointer(&s, []int{6, 7, 8, 9})
	fmt.Println(s)
}
