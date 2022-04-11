package main

import "fmt"

type Node struct {
	Val int
	Next *Node
}

func Add(h **Node, val int) {
	// fmt.Println("here")
	n := Node{Val:val,Next:nil}

	if *h==nil {
		fmt.Println("here added ",n.Val)

		*h=&n
		return
	}
	fmt.Println("2here head.val ",(*h).Val)

	p := *h
	for (p).Next != nil {
		(p)=(p).Next
		fmt.Printf("i ")
	}
	fmt.Println()
	(p).Next = &n

}

func PrintList(h *Node) {
	if h==nil {
		fmt.Println("empty list")
		return
	}
	// p := h
	for h!=nil {
		fmt.Printf("%d ",(h).Val)
		if h==nil {
			fmt.Println("it's still empty, LOL")
		}
		h=(h).Next
	}
	fmt.Println()
}


func main() {
	var head  *Node
	head = nil

	for i:=0; i<10; i++ {
		Add(&head,i+1)
		if head==nil {
			fmt.Println("still empty!! WTF")
		}
	}
	fmt.Println(head.Val)
	PrintList(head)
}