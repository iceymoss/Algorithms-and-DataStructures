package mid

type ListNode struct {
	Val  int
	Next *ListNode
}

// 思路：不一定需要删除真的节点，我们拿不到node的父亲节点，从逻辑上我们无法删除node
// 所以我们只需要将node的后的节点的值赋值给node的值，然后将node.Next删除即可
func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
