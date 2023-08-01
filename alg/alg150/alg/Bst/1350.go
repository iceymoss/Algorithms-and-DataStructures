package Bst

import "sort"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	ans1, ans2 := make([]int, 0), make([]int, 0)
	dfs(root1, &ans1)
	dfs(root2, &ans2)
	ans1 = append(ans1, ans2...)
	sort.Slice(ans1, func(i, j int) bool {
		return ans1[i] < ans1[j]
	})
	return ans1
}

func dfs(root *TreeNode, ans *[]int) {
	if root == nil {
		return
	}
	dfs(root.Left, ans)
	*ans = append(*ans, root.Val)
	dfs(root.Right, ans)
}

func bfs(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	ans := make([]int, 0)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		size := len(queue)
		for size != 0 {
			top := queue[0]
			queue = queue[1:]
			ans = append(ans, top.Val)
			if top.Left != nil {
				queue = append(queue, top.Left)
			}
			if top.Right != nil {
				queue = append(queue, top.Right)
			}
			size--
		}
	}
	return ans
}
