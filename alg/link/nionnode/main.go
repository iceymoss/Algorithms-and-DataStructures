package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 使用哈希表用来存储节点
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	hashesMap := make(map[*ListNode]bool)
	for tmp := headA; tmp != nil; tmp = tmp.Next {
		hashesMap[tmp] = true
	}

	for tmp := headB; tmp != nil; tmp = tmp.Next {
		if hashesMap[tmp] {
			return tmp
		}
	}
	return nil
}

func linkLeftToRight(link *ListNode) *ListNode {
	if link == nil {
		return nil
	}
	cur := link
	var per *ListNode
	for cur != nil {
		next := cur.Next
		cur.Next = per
		per = cur
		cur = next
	}
	return per
}

func main() {

}
