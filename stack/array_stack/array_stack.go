package array_stack

import "sync"

type ArrayStack struct {
	Array []int
	Size  int
	Lock  sync.Mutex
}

//Push 入栈
func (s *ArrayStack) Push(element int) {
	s.Lock.Lock()
	defer s.Lock.Unlock()

	s.Array = append(s.Array, element)
	s.Size++
}

//Pop 出栈
func (s *ArrayStack) Pop() int {
	s.Lock.Lock()
	defer s.Lock.Unlock()

	if s.Size == 0 {
		panic("empty")
	}
	top := s.Array[s.Size-1]

	newArr := make([]int, s.Size-1, s.Size-1)
	for index := 0; index < s.Size-1; index++ {

		newArr[index] = s.Array[index]
	}

	s.Array = newArr
	s.Size--
	return top
}

func (s *ArrayStack) Peek() int {
	if s.Size == 0 {
		panic("empty")
	}
	return s.Array[s.Size-1]
}

func (s *ArrayStack) IsEmpty() bool {
	return s.Size == 0
}

func (s *ArrayStack) Len() int {
	return s.Size
}
