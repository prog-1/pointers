package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

func (n *Node) asSlice() (slice []int) {
	if n == nil {
		return nil
	}

	for n.Next != nil {
		//fmt.Println(n.Next)
		slice = append(slice, n.Value)
		n = n.Next
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
	a := *head
	if a == nil {
		*head = &Node{Value: x, Next: nil}
		return
	}
	for a.Next != nil {
		a = a.Next
	}
	a.Next = &Node{Value: x, Next: nil}

}

func Insert(n **Node, x int) {
	// if *n == nil {
	// 	*n = &Node{Value: x, Next: nil}
	// }
	a := *n
	var prev *Node
	for a.Value < x && a.Next != nil {
		prev, a = a, a.Next
	}
	if prev == nil {
		Append(n, x)
		return
	}
	prev.Next = &Node{Value: x, Next: a}
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
	// fmt.Println(*n == nil)
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
	for _, v := range []int{1, 2, 4, 5} {
		Append(&head, v)
	}
	Insert(&head, 3)
	fmt.Println(head.asSlice())
}
