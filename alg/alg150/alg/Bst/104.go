package Bst

func maxDepth(root *TreeNode) int {
	//深度优先遍历
	if root == nil {
		return 0
	}
	//记录当前层
	layerLeft := maxDepth(root.Left)
	layerRight := maxDepth(root.Right)
	if layerRight > layerLeft {
		return layerRight + 1
	}
	if layerRight < layerLeft {
		return layerLeft + 1
	}
	return layerRight + 1
}
