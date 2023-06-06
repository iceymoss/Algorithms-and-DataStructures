package Bst

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func goodNodes(root *TreeNode) int {
	res := 1
	dfs(root, &res, root.Val)
	return res
}

func dfs(root *TreeNode, count *int, maxTag int) {
	if (root.Left == nil && root.Right == nil) || root == nil {
		return
	}

	if root.Left != nil {
		lmax := max(maxTag, root.Left.Val)
		if root.Left.Val >= lmax {
			*count++
		}
		dfs(root.Left, count, lmax)
	}

	if root.Right != nil {
		rmax := max(maxTag, root.Right.Val)
		if root.Right.Val >= rmax {
			*count++
		}
		dfs(root.Right, count, rmax)
	}
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
