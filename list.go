package main

import "fmt"

type Node struct {
	Val int
	Next *Node
}

func AddLast(h **Node, val int) {
	n := Node{Val:val,Next:nil}

	if *h==nil {
		*h=&n
		return
	}
	p := *h
	for (p).Next != nil {
		(p)=(p).Next
	}
	(p).Next = &n

}
func AddFront(h **Node, val int) {
	n := Node{Val:val,Next:nil}
	if *h == nil {
		*h = &n
		return
	}
	n.Next = *h
	(*h) = &n
}

func PrintList(h *Node) {
	if h==nil {
		return
	}
	// p := h
	for h!=nil {
		fmt.Printf("%d ",(h).Val)
		h=(h).Next
	}
	fmt.Println()
}


func main() {
	var head  *Node
	head = nil

	for i:=0; i<10; i++ {
		AddFront(&head,i+1)
	}
	PrintList(head)
}