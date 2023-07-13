package Link

type ListNodes struct {
	Val  int
	Next *ListNodes
}

// 直接将链表的元素进行拼接成完整的数字然后相加，再拆分
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1.Val+l2.Val == 0 {
		return l1
	}
	//遍历l1
	cur1 := l1
	cur2 := l2
	var tage1 int64
	for cur1 != nil {
		tage1 += int64(cur1.Val)
		cur1 = cur1.Next
		if cur1 != nil {
			tage1 *= 10
		}
	}
	var tage2 int64
	for cur2 != nil {
		tage2 += int64(cur2.Val)
		cur2 = cur2.Next
		if cur2 != nil {
			tage2 *= 10
		}
	}
	total := tage1 + tage2
	//使用栈
	stack := make([]int64, 0)
	for total != 0 {
		num := total % 10
		//入栈
		stack = append(stack, num)
		total /= 10
	}
	head := &ListNode{}
	cur := head
	for i := len(stack) - 1; i >= 0; i-- {
		next := &ListNode{Val: int(stack[i])}
		cur.Next = next
		cur = next
	}
	return head.Next
}

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	var s1, s2 []int

	//入栈
	for l1 != nil {
		s1 = append(s1, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		s2 = append(s2, l2.Val)
		l2 = l2.Next
	}

	//栈先进后出，我们可以从个位数相加然后将结果构建链表节点, curry用于向高位推进，例如：sum = 5, curry = 40, sum += curry == 45
	curry := 0
	head := &ListNode{}
	for len(s1) > 0 || len(s2) > 0 || curry > 0 {
		//移除栈顶
		sum := 0
		if len(s1) > 0 {
			sum += s1[len(s1)-1]
			s1 = s1[:len(s1)-1]
		}
		if len(s2) > 0 {
			sum += s2[len(s2)-1]
			s2 = s2[:len(s2)-1]
		}

		sum += curry
		//构造节点
		node := &ListNode{Val: sum % 10}
		//链表节点逆向构建
		node.Next = head.Next
		head.Next = node
		//向高位推进
		curry = sum / 10
	}
	return head.Next
}
