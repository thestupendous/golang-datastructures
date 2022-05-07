package main

import "fmt"

var list []int

type tree struct {
	data  int
	left  *tree
	right *tree
}

func isBst(head *tree) {
	if head.left != nil {
		isBst(head.left)
	}
	list = append(list, head.data)
	if head.right != nil {
		isBst(head.right)
	}
}

func main() {
	head := &tree{83, nil, nil}
	head.left = &tree{8, nil, nil}
	head.right = &tree{983, nil, nil}

	isBst(head)
	if len(list) < 2 {
		return
	}
	yes := true
	prev, cur := list[0], list[1]
	for i := 1; i < len(list); i++ {
		cur = list[i]
		if prev > cur {
			yes = false
		}
		prev = cur
	}
	fmt.Println(list)

	if yes {
		fmt.Println("yes it is bst!!!")
	} else {
		fmt.Println("NOOOOOOOOO")
	}

}
