package Link

type ListNode struct {
	Val  int
	Next *ListNode
}

func DeleteMiddle(head *ListNode) *ListNode {
	//1、遍历列表，找出n的值
	count := 0
	tag := 0
	cur := head
	for cur != nil {
		cur = cur.Next
		count++
	}

	tag = count / 2
	cur = head
	per := cur
	//移除对应位置的节点
	for i := 0; i <= tag && cur != nil; {
		if count <= 1 {
			return nil
		}
		if i < tag {
			//记录前序节点
			per = cur
		}
		if i == tag {
			per.Next = cur.Next
		}
		cur = cur.Next
		i++
	}
	return head
}
