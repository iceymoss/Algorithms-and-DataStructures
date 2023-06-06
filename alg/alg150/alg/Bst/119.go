package Bst

// 使用广度优先遍历，依次遍历二叉树的每一层节点，然后找到队列最后一个节点值加入返回值即可
func rightSideView(root *TreeNode) []int {
	res := make([]int, 0)
	if root == nil {
		return res
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		laySize := len(queue)
		for laySize != 0 {
			cur := queue[0]
			if laySize == 1 {
				res = append(res, cur.Val)
			}
			queue = queue[1:]
			if cur.Left != nil {
				queue = append(queue, cur.Left)
			}
			if cur.Right != nil {
				queue = append(queue, cur.Right)
			}
			laySize--
		}
	}
	return res
}
