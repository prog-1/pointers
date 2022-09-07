package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

func Prepend(head **Node, x int) {
	*head = &Node{x, *head}
}

func Append(head **Node, x int) {
	if *head == nil {
		*head = &Node{x, nil}
	} else {
		p := *head // I don't know why, but it doesn't work otherwise
		for p.Next != nil {
			p = p.Next
		}
		p.Next = &Node{x, nil}
	}
}

//func Insert(head **Node, x int)

//func Contains(head *Node, x int) bool

//func Remove(head **Node, x int) (removed bool)

func main() {
	var head *Node
	Append(&head, 7)
	Prepend(&head, 3)
	//Insert(&head, 4)
	//Insert(&head, 2)
	//Insert(&head, 9)
	//Remove(&head, 4)
	for cur := head; cur != nil; cur = cur.Next {
		fmt.Print(cur.Value, " ")
	} // Output: 2 3 7 9 (should be)
	fmt.Println() // Current output: 3 7
}
