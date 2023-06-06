package Bst

func deleteNode(root *TreeNode, key int) *TreeNode {
	//使用递归的方法
	//1、先找到对应的key的位置
	if root == nil {
		return nil
	} else if root.Val > key {
		root.Left = deleteNode(root.Left, key)

	} else if root.Val < key {
		root.Right = deleteNode(root.Right, key)

	} else if root.Left == nil || root.Right == nil { //存在左子树或者右子树,或者该节点为叶子节点
		if root.Left != nil {
			return root.Left
		}
		return root.Right
	} else { //存在左右子树
		successor := root.Right
		//找到右子树的最小节点
		for successor.Left != nil {
			successor = successor.Left
		}
		//这个过程我们还需要删除右子树的最小节点
		successor.Right = deleteNode(root.Right, successor.Val)
		successor.Left = root.Left
		return successor
	}
	return root
}
