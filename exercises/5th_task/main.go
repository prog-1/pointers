package main

type Node struct {
	Value int
	Next  *Node
}

func (n Node) asSlice() (slice []int) {
	i := 0
	for n.Next != nil {
		//fmt.Println(n.Next)
		slice = append(slice, n.Value)
		n = *n.Next
		i++
	}
	slice = append(slice, n.Value)
	return
}

func Prepend(head **Node, x int) {
	*head = &Node{Value: x, Next: *head}
}

func Append(head **Node, x int) {
	a := *head
	for a.Next != nil {
		a = a.Next
	}
	a.Next = &Node{Value: x, Next: nil}

}

func Insert(head **Node, x int) {
	a := *head
	var prev *Node
	for a.Value < x && a.Next != nil {
		prev, a = a, a.Next
	}
	if prev == nil {
		Append(head, x)
		return
	}
	prev.Next = &Node{Value: x, Next: a}
}

func Contains(head *Node, x int) bool {
	a := head
	for a.Next != nil {
		if a.Value == x {
			return true
		}
		a = a.Next
	}
	return a.Value == x
}

func Remove(head **Node, x int) (removed bool) {
	a := *head
	var prev *Node
	for a.Next != nil {
		if a.Value == x {
			if prev != nil {
				prev.Next = a.Next
			} else {
				*head = a.Next
			}
			return true
		}
		prev, a = a, a.Next
	}
	if a.Value == x {
		prev.Next = a.Next
		return true
	}
	return false
}

func main() {

}
