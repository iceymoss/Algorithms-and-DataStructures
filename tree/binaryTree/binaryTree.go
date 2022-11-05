package binaryTree

import (
	"fmt"
)

//TreeNode 构造二叉树节点
type TreeNode struct {
	Data  int
	Left  *TreeNode
	Right *TreeNode
}

//PreOrder 前序遍历
func PreOrder(tree *TreeNode) {
	if tree == nil {
		return
	}

	fmt.Println(tree.Data)
	PreOrder(tree.Left)
	PreOrder(tree.Right)
}

//MidOrder 中序遍历
func MidOrder(tree *TreeNode) {
	if tree == nil {
		return
	}

	MidOrder(tree.Left)
	fmt.Println(tree.Data)
	MidOrder(tree.Right)
}

//PostOrder 后序遍历
func PostOrder(tree *TreeNode) {
	if tree == nil {
		return
	}

	MidOrder(tree.Left)
	MidOrder(tree.Right)
	fmt.Println(tree.Data)

}

//LayerOrder 层序遍历
func LayerOrder(tree *TreeNode) {
	if tree == nil {
		return
	}

	//辅助队列
	queue := LinkQueue{}
	queue.Push(tree)

	for !queue.IsEmpty() {

		t := queue.Pop()

		fmt.Println(t.Data)

		if t.Left != nil {
			queue.Push(t.Left)
		}
		if t.Right != nil {
			queue.Push(t.Right)
		}
	}

}
