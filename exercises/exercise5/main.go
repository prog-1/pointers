package main

import (
	"fmt"
)

type Node struct {
	Value int
	Next  *Node
}

func Prepend(head **Node, x int) {

	*head = &Node{Value: x, Next: *head}
}

func Append(head **Node, x int) {

	current := *head
	for current.Next != nil {
		current = current.Next
	}

	Prepend(&current, x)
}

func Insert(head **Node, x int) {

	if x < (*head).Value {
		Prepend(head, x)
		return
	}

	current := *head
	for current.Next != nil && current.Next.Value < x {
		current = current.Next
	}

	inserted := &Node{Value: x, Next: current.Next}
	current.Next = inserted

}

func Contains(head *Node, x int) bool {

	current := head
	for current.Next != nil {
		if current.Value == x {
			return true
		}
	}

	return false
}

func Remove(head **Node, x int) (removed bool) {

	current := (*head)

	if current.Value == x {
		(*head) = current.Next
		return true
	}

	for current.Next != nil {
		if current.Next.Value == x {
			// Deleting Node
			current.Next = current.Next.Next
			return true
		}

		current = current.Next
	}
	
	return false
}

func main() {
	head := &Node{2, nil}

	Append(&head, 7)
	Prepend(&head, 3)
	Insert(&head, 0)

	for cur := head; cur != nil; cur = cur.Next {
		fmt.Print(cur.Value, " ")
	}
	fmt.Println()

	fmt.Println(Remove(&head, 2))
	for cur := head; cur != nil; cur = cur.Next {
		fmt.Print(cur.Value, " ")
	}
}
