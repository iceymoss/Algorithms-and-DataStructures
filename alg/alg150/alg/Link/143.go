package Link

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// func reorderList(head *ListNode)  {
//     //转为线性表，使用双指针法
//     arr := make([]*ListNode, 0)
//     cur := head
//     for cur != nil {
//         arr = append(arr, cur)
//         cur = cur.Next
//     }
//     l, r := 0, len(arr)-1
//     for l < r {
//        arr[l].Next = arr[r]
//        l++
//        if l == r {
//            break
//        }
//        arr[r].Next = arr[l]
//        r--
//     }
//     arr[l].Next = nil
// }

// 获取链表中间件节点
func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// 反转链表
func reverseList(head *ListNode) *ListNode {
	var per, cur *ListNode = nil, head
	for cur != nil {
		Next := cur.Next
		cur.Next = per
		per = cur
		cur = Next
	}
	return per
}

// 双指针+反转链表+合并链表
func reorderList(head *ListNode) {
	mid := middleNode(head)
	head2 := reverseList(mid)
	for head2.Next != nil {
		next := head.Next
		next2 := head2.Next
		head.Next = head2
		head2.Next = next
		head = next
		head2 = next2
	}
}
