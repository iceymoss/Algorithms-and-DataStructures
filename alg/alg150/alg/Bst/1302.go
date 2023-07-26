package Bst

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deepestLeavesSum(root *TreeNode) int {
	//广度优先遍历
	res := 0
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		//queue_size控制层
		queue_size := len(queue)
		total := 0
		for queue_size != 0 {
			cur := queue[0]
			total += cur.Val
			queue = queue[1:]
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
			queue_size--
		}
		res = total
	}
	return res
}

// 深度优先遍历
func deepestLeavesSumDfs(root *TreeNode) int {
	maxLevel := -1
	sum := 0
	var dfs func(*TreeNode, int)
	dfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}
		if level > maxLevel {
			maxLevel = level
			sum = root.Val
		} else if level == maxLevel {
			sum += root.Val
		}
		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}
	dfs(root, 0)
	return sum
}
