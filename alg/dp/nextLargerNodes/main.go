package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

// leetCode 1019
// 核心逻辑：判断链表后续相邻节点值，是否大于当前节点，如果大于则将后续相邻节点值加入返回结果中，如果小于向下遍历，并且需要记录当前节点，找到一个更大的值
// 本题需要使用动态规划和单调栈(递增或者递减)本题递减，需要自底向上进行操作，从链表最后节点往前遍历只需要到比当前遍历到的所有节点中最大的节点还要大，即可
// 所以我们先反转链表，从而实现自底向上，当然需要使用栈维护一个递减的状态。
func nextLargerNodes(head *ListNode) []int {
	//自底向上，所以需要反转链表
	list, n := rightToLeft(head)
	ans := make([]int, n)
	stack := make([]int, 0)

	//遍历反转后的链表
	for cur := list; cur != nil; cur = cur.Next {
		//维护单调栈， 只要栈顶元素小于当前节点值，就将栈顶元素弹出
		for len(stack) > 0 && stack[len(stack)-1] <= cur.Val {
			stack = stack[:len(stack)-1]
		}

		n--
		//栈顶元素大于当前节点值，需要将栈顶元素放入结果中
		if len(stack) > 0 {
			ans[n] = stack[len(stack)-1]
		}
		//将当前节点值入栈
		stack = append(stack, cur.Val)
	}
	return ans
}

// 反转链表
func rightToLeft(head *ListNode) (*ListNode, int) {
	cur := head
	var per *ListNode
	n := 0
	for cur.Next != nil {
		next := cur.Next
		cur.Next = per
		per = cur
		cur = next
		n++
	}
	cur.Next = per
	return cur, n + 1
}

func main() {
	list := &ListNode{2, &ListNode{1, &ListNode{5, nil}}}
	fmt.Println(nextLargerNodes(list))

}
