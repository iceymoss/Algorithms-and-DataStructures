package Bst

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func MidDfs(root *Node) {
	if root == nil {
		return
	}

	MidDfs(root.Left)
	fmt.Println(root.Val)
	MidDfs(root.Right)
}

func Bfs(root *Node) {
	if root == nil {
		return
	}

	queue := make([]*Node, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		laySize := len(queue)
		for laySize != 0 {
			curNode := queue[0]
			queue = queue[1:]
			if curNode.Left != nil {
				queue = append(queue, curNode.Left)
			}

			if curNode.Right != nil {
				queue = append(queue, curNode.Right)
			}
			laySize--
		}
	}

}
