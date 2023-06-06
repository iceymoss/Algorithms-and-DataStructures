package Bst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//解题思路：使用深度优先遍历遍历整个树，当遍历到子节点时将子节点的值加入到序列中，然后比较两个序列

func LeafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	if root2 == nil && root1 == nil {
		return false
	}
	//使用中序遍历
	arr1, arr2 := make([]int, 0), make([]int, 0)
	traver(root1, &arr1)
	traver(root2, &arr2)

	if len(arr1) != len(arr2) {
		return false
	}

	for i := 0; i < len(arr2); i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}

func traver(root *TreeNode, arr *[]int) {
	if root != nil && root.Right == nil && root.Left == nil {
		*arr = append(*arr, root.Val)
		return
	}
	if root != nil {
		traver(root.Left, arr)
		traver(root.Right, arr)
	}
}
