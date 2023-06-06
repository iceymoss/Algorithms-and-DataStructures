package Bst

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	return dfsBST(root, val)
}

func dfsBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}

	if root.Val > val {
		return dfsBST(root.Left, val)
	}
	if root.Val < val {
		return dfsBST(root.Right, val)
	}

	return nil
}
