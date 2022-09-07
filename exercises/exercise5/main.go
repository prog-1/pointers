package main

import (
	"fmt"
)

type Node struct {
	Value int
	Next  *Node
}

func Prepend(head **Node, x int) {
	// Plan:
	// * Create new head
	// * Make current head new head's next
	// * Addign new head to current head pointer

	var newHead *Node = new(Node)
	newHead.Value = x
	newHead.Next = *head

	*head = newHead
}

func Append(head **Node, x int) {

	var current *Node = *head
	for current.Next != nil {
		current = current.Next
	}

	var appended *Node = new(Node)
	appended.Value = x
	appended.Next = nil
	current.Next = appended
}

func main() {
	var head *Node = new(Node)
	head.Value = 0

	Prepend(&head, 3)
	Append(&head, 7)

	for cur := head; cur != nil; cur = cur.Next {
		fmt.Print(cur.Value, " ")
	}
	fmt.Println()

}
