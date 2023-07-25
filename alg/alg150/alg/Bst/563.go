package Bst

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTilt(root *TreeNode) int {
	res := 0
	getvalue(root, &res)
	return res
}

func getvalue(root *TreeNode, res *int) int {
	if root == nil {
		return 0
	}
	l := getvalue(root.Left, res)
	r := getvalue(root.Right, res)
	*res += abs(l, r)
	return l + r + root.Val
}

func abs(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}
