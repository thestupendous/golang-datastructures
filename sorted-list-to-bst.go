package main

import "fmt"

type tree struct {
	data  int
	left  *tree
	right *tree
}

func listToBst(list []int, start int, end int) *tree {
	if start > end || start < 0 || end >= len(list) {
		return nil
	}
	node := tree{0, nil, nil}
	if start == end {
		node.data = list[start]
		return &node
	}
	mid := (start + end) / 2

	node.data = list[mid]

	//left subtree
	node.left = listToBst(list, start, mid-1)

	//right subtree
	node.right = listToBst(list, mid+1, end)

	return &node

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
	// head := &tree{83, nil, nil}
	// head.left = &tree{8, nil, nil}
	// head.right = &tree{983, nil, nil}

	arr := []int{2, 19, 23, 41, 50, 231, 5012, 201321}

	head := listToBst(arr, 0, len(arr)-1)
	fmt.Println(arr)
	inorder(head)
	fmt.Println()

	// if yes {
	// fmt.Println("yes it is bst!!!")
	// } else {
	// fmt.Println("NOOOOOOOOO")
	// }

}
