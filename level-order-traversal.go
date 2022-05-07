package main

import "fmt"

type tree struct {
	data  int
	left  *tree
	right *tree
}

func levelOrder(list []int, lArray *[]int, start int, end int, count int) *tree {
	if start > end || start < 0 || end >= len(list) {
		return nil
	}
	count++
	node := tree{0, nil, nil}
	if start == end {
		node.data = list[start]
		*lArray = append(*lArray, count)
		return &node
	}
	mid := (start + end) / 2

	node.data = list[mid]
	*lArray = append(*lArray, count)

	//left subtree
	node.left = levelOrder(list,lArray, start, mid-1, count)

	//right subtree
	node.right = levelOrder(list, lArray, mid+1, end, count)

	return &node

}

func inorder(head *tree) {
	if head == nil {
		return
	}

	inorder(head.left)
	fmt.Printf("%d ", head.data)
	inorder(head.right)
	fmt.Println(

}

func main() {
	// head := &tree{83, nil, nil}
	// head.left = &tree{8, nil, nil}
	// head.right = &tree{983, nil, nil}

	arr := []int{2, 19, 23, 41, 50, 231, 5012, 201321}
	var levelArr []int

	head := levelOrder(arr, &levelArr, 0, len(arr)-1, 0)
	fmt.Println(arr)
	inorder(head)
	fmt.Println()

	// if yes {
	// fmt.Println("yes it is bst!!!")
	// } else {
	// fmt.Println("NOOOOOOOOO")
	// }

}
