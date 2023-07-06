package Bst

import "strconv"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func binaryTreePaths(root *TreeNode) []string {
	//递归结束条件
	res := make([]string, 0)
	if root == nil {
		return res
	}
	node(root, &res, "")
	return res
}

func node(root *TreeNode, paths *[]string, path string) {
	if root != nil {
		pathSB := path
		pathSB += strconv.Itoa(root.Val)
		if root.Left == nil && root.Right == nil {
			*paths = append(*paths, pathSB)
		} else {
			pathSB += "->"
			node(root.Left, paths, pathSB)
			node(root.Right, paths, pathSB)
		}
	}
}
