package link_stack

import "sync"

//LinkNode 构造链表节点
type LinkNode struct {
	Next  *LinkNode
	Value int
}

//LinkStack 构造栈
type LinkStack struct {
	Root *LinkNode
	Size int
	Lock sync.Mutex
}

//Push 入栈
func (s *LinkStack) Push(value int) {
	s.Lock.Lock()
	defer s.Lock.Unlock()

	//栈顶为空
	if s.Root == nil {
		newNode := new(LinkNode)
		newNode.Value = value
		s.Root = newNode
	} else {
		sta := s.Root

		newNode := new(LinkNode)
		newNode.Value = value

		newNode.Next = sta

		s.Root = newNode
	}
	s.Size++
}

func (s *LinkStack) Pop() int {
	s.Lock.Lock()
	defer s.Lock.Unlock()

	if s.Size == 0 {
		panic("empty")
	}

	top := s.Root

	s.Root = s.Root.Next
	s.Size--
	return top.Value
}

//Peek 获取栈顶
func (s *LinkStack) Peek() int {
	if s.Size == 0 {
		panic("empty")
	}
	return s.Root.Value
}

//IsEmpty 栈是否为空
func (s *LinkStack) IsEmpty() bool {
	return s.Root == nil
}

//Len 返回栈深度
func (s *LinkStack) Len() int {
	return s.Size
}
