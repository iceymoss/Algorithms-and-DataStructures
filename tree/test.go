package main

import (
	"Algorithms-and-DataStructures/tree/binaryTree"
)

func main() {
	tree := &binaryTree.TreeNode{Data: 1}
	tree.Left = &binaryTree.TreeNode{Data: 2}
	tree.Left.Left = &binaryTree.TreeNode{Data: 3}
	tree.Left.Right = &binaryTree.TreeNode{Data: 4}

	t := &binaryTree.TreeNode{Data: 100}
	tree.Insert(t)
	binaryTree.LayerOrder(tree)
}
