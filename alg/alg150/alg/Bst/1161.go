package Bst

import "math"

// 广度优先遍历，计算每一层的和进行比较，找出最大值，并返回对应层数
func maxLevelSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	resLay := 1
	maxV := math.MinInt
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	lay := 1
	for len(queue) != 0 {
		size := len(queue)
		total := 0
		for size != 0 {
			cur := queue[0]
			queue = queue[1:]
			total += cur.Val
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
			size--
		}
		if total > maxV {
			resLay = lay
			maxV = total
		}
		lay++
	}
	return resLay
}
