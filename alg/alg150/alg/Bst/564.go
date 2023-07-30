package Bst

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func constructMaximumBinaryTree(nums []int) *TreeNode {
	root := &TreeNode{}
	buildTree(nums, root)
	return root
}

func buildTree(nums []int, cur *TreeNode) {
	if len(nums) == 0 {
		return
	}
	i := getMaxValue(nums)
	cur.Val = nums[i]
	if len(nums[0:i]) != 0 {
		cur.Left = &TreeNode{}
	}
	if len(nums[i+1:]) != 0 {
		cur.Right = &TreeNode{}
	}
	buildTree(nums[0:i], cur.Left)
	buildTree(nums[i+1:], cur.Right)
}

func getMaxValue(nums []int) int {
	max := math.MinInt
	tag := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > max {
			tag = i
		}
	}
	return tag
}
