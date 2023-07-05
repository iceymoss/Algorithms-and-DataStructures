package Bst

type TreeNodes struct {
	Val   int
	Left  *TreeNodes
	Right *TreeNodes
}

// 是指针，直接递归到底，进行交换节点即可
func invertTree(root *TreeNodes) *TreeNodes {
	r := root
	invert(root)
	return r
}

func invert(root *TreeNodes) {
	//递归终止条件
	if root == nil {
		return
	}
	root.Left, root.Right = root.Right, root.Left
	invert(root.Left)
	invert(root.Right)
}
