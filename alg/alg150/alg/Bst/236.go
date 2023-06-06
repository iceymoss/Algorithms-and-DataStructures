package Bst

// 使用回溯法
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	return findNode(root, p, q)
}

func findNode(curNode, p, q *TreeNode) *TreeNode {
	//递归终止条件, 如果当前节点为空，或者找到了q或者p
	if curNode == nil || curNode == p || curNode == q {
		return curNode
	}

	//进入下一层
	left := findNode(curNode.Left, p, q)
	right := findNode(curNode.Right, p, q)

	//并回溯,处理回溯结果
	if left == nil {
		return right
	}

	if right == nil {
		return left
	}

	//如果都为空，则返回当前节点
	return curNode
}
