package binaryTree

import "sync"

//LinkNode 构造节点
type LinkNode struct {
	Next  *LinkNode
	Value *TreeNode
}

//LinkQueue 构造队列
type LinkQueue struct {
	Root *LinkNode
	Size int
	Lock sync.Mutex
}

//Push 入队
func (q *LinkQueue) Push(v *TreeNode) {
	q.Lock.Lock()
	defer q.Lock.Unlock()

	if q.Root == nil {
		node := new(LinkNode)
		node.Value = v
		q.Root = node
	} else {
		current := q.Root
		for {
			if current.Next == nil {
				node := new(LinkNode)
				node.Value = v
				current.Next = node
				break
			}
			current = current.Next
		}
	}
	q.Size++
}

func (q *LinkQueue) Pop() *TreeNode {
	q.Lock.Lock()
	defer q.Lock.Unlock()

	if q.Size == 0 {
		panic("empty")
	}

	first := q.Root

	q.Root = q.Root.Next

	q.Size--
	return first.Value
}

func (q *LinkQueue) Peek() *TreeNode {
	if q.Size == 0 {
		panic("empty")
	}

	first := q.Root
	return first.Value
}

func (q *LinkQueue) IsEmpty() bool {
	return q.Size == 0
}

func (q *LinkQueue) Len() int {
	return q.Size
}
