package main

import (
	"fmt"
	"log"
	"strconv"
)

func Comp(node1, node2 Tree) byte { // n1<n2 -1 ; n1>n2 +1 ; n1==n2 0
	num1, err1 := strconv.Atoi(node1.Data)
	num2, err2 := strconv.Atoi(node2.Data)
	_ = err1
	_ = err2
	if num1 < num2 {
		return 5
	} else if num1 == num2 {
		return 10
	} else {
		return 15
	}
}

type Tree struct {
	Data  string
	Left  *Tree
	Right *Tree
}

func InitTree(node *Tree, value string) {
	node.Left = nil
	node.Right = nil
	node.Data = value
}

func Insert(head **Tree, value string) {
	var node Tree
	InitTree(&node, value)
	if *head == nil {
		*head = &node
		log.Println("added first node")
		return
	}
	ptr := *head
	for ptr.Left != nil || ptr.Right != nil {
		if Comp(node, *ptr) < 10 {
			if ptr.Left != nil {
				ptr = ptr.Left
			} else {
				ptr.Left = &node
				log.Println("added one more node to tree")
				return
			}
		} else if Comp(node, *ptr) > 10 {
			if ptr.Right != nil {
				ptr = ptr.Right
			} else {
				ptr.Right = &node
				log.Println("added one more node to tree")
				return
			}
		} else {
			log.Fatal("ERROR: requested value alredy matches with existing tree node's value!")
			return
		}
	}
	log.Fatal("ERROR: Are you time traveller? How'd you reach here")
}

func Traverse(head *Tree) {
	if head.Left != nil {
		Traverse(head.Left)
	}
	log.Printf("[%s] ", head.Data)
	if head.Right != nil {
		Traverse(head.Right)
	}
}

func main() {
	var head *Tree

	var n int
	//	var arr []string
	var s string
	fmt.Println("enter size of tree : ")
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%s", &s)
		Insert(&head, s)
		//	arr = append(arr, s)
	}
	Traverse(head)

}
