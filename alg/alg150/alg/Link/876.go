package Link

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func middleNode(head *ListNode) *ListNode {
	count := 0
	cur := head
	for cur != nil {
		cur = cur.Next
		count++
	}

	tag := count / 2
	cur = head
	res := cur
	for i := 0; i <= tag && cur != nil; i++ {
		if i == tag {
			res = cur
		}
		cur = cur.Next
	}
	return res
}

func middleNode1(head *ListNode) *ListNode {
	low, fast := head, head
	for fast != nil && fast.Next != nil {
		low = low.Next
		fast = fast.Next.Next
	}
	return low
}
