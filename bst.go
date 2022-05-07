package main

import "fmt"

var list []int

type tree struct {
	data  int
	left  *tree
	right *tree
}

func inorder(head *tree) {
	if head == nil {
		return
	}

	inorder(head.left)
	fmt.Printf("%d ", head.data)
	inorder(head.right)

}

func main() {
	head := &tree{83, nil, nil}
	head.left = &tree{8, nil, nil}
	head.right = &tree{983, nil, nil}

}
