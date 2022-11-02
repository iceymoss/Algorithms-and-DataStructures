package main

import (
	"Algorithms-and-DataStructures/queue/array_queue"
	"Algorithms-and-DataStructures/queue/link_queue"
	"fmt"
)

func TestArrayQueue() {
	q := new(array_queue.ArrayQueue)
	for i := 10; i > 0; i-- {
		q.Push(i)
	}

	fmt.Println("size:", q.Size)

	for i := 0; i < 10; i++ {
		fmt.Println(q.Pop())
	}

	fmt.Println(q.IsEmpty())
}

func TestLinkQueue() {
	q := new(link_queue.LinkQueue)
	for i := 10; i > 0; i-- {
		q.Push(i)
	}

	fmt.Println("size:", q.Size)

	for i := 0; i < 10; i++ {
		fmt.Println(q.Pop())
	}

	fmt.Println(q.IsEmpty())
}

func main() {
	TestArrayQueue()
	TestLinkQueue()
}
