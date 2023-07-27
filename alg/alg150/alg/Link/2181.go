package Link

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeNodes(head *ListNode) *ListNode {
	//直接遍历累加
	sum := 0
	per := head
	cur := head.Next
	for cur != nil {
		if cur.Val == 0 {
			cur.Val = sum
			per.Next = cur
			cur = cur.Next
			per = per.Next
			sum = 0
		} else {
			sum += cur.Val
			cur = cur.Next
		}
	}
	return head.Next
}
