package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

func (n Node) asSlice() (slice []int) {
	for n.Next != nil {
		//fmt.Println(n.Next)
		slice = append(slice, n.Value)
		n = *n.Next
	}
	slice = append(slice, n.Value)
	return
}

func Prepend(head **Node, x int) {
	if head == nil {
		*head = &Node{Value: x, Next: nil}
		return
	}
	*head = &Node{Value: x, Next: *head}
}

func Append(head **Node, x int) {
	if head == nil {
		*head = &Node{Value: x, Next: nil}
		return
	}
	a := *head
	for a.Next != nil {
		a = a.Next
	}
	a.Next = &Node{Value: x, Next: nil}

}

func Insert(n **Node, x int) {
	var prev *Node
	for *n != nil && (*n).Value < x {
		prev, n = *n, &((*n).Next)
	}
	if prev != nil {
		prev.Next = &Node{Value: x, Next: *n}
		return
	}
	*n = &Node{Value: x, Next: *n}
}

func Contains(head *Node, x int) bool {
	if head == nil {
		return false
	}
	a := head
	for a.Next != nil {
		if a.Value == x {
			return true
		}
		a = a.Next
	}
	return a.Value == x
}

func Remove(n **Node, x int) (removed bool) {
	for *n != nil && (*n).Value != x {
		n = &((*n).Next)
	}
	if *n != nil {
		*n = (*n).Next
		return true
	}
	return false
}

func main() {
	var head *Node
	fmt.Println(Contains(head, 1))
}
