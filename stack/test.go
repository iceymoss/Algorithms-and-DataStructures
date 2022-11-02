package main

import (
	"Algorithms-and-DataStructures/stack/array_stack"
	"Algorithms-and-DataStructures/stack/link_stack"
	"fmt"
)

func TestLinkStack() {
	s := new(link_stack.LinkStack)

	arr := []int{1, 23, 4, 45}
	for _, v := range arr {
		s.Push(v)
	}

	fmt.Println("size", s.Size)

	len := s.Size
	for i := 0; i < len; i++ {
		fmt.Println(s.Pop())
	}
}

func TestArrayStack() {
	s := new(array_stack.ArrayStack)

	arr := []int{1, 23, 4, 45}
	for _, v := range arr {
		s.Push(v)
	}

	fmt.Println("size", s.Size)

	len := s.Size
	for i := 0; i < len; i++ {
		fmt.Println(s.Pop())
	}

}

func main() {
	TestArrayStack()
	TestLinkStack()
}
