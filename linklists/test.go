package main

import (
	"Algorithms-and-DataStructures/linklists/doubleList"
	"fmt"
)

func main() {
	// 在列表头部插入新元素
	list := new(doubleList.DoubleList)
	list.AddNodeFromHead(0, "I")
	list.AddNodeFromHead(0, "love")
	list.AddNodeFromHead(0, "you")
	// 在列表尾部插入新元素
	list.AddNodeFromHead(0, "哈哈哈")
	list.AddNodeFromHead(0, "淦饭")

	// 正常遍历，比较慢，因为内部会遍历拿到值返回
	for i := 0; i < list.GetLen(); i++ {
		// 从头部开始索引
		node := list.GetFormHead(i)

		// 节点为空不可能，因为list.Len()使得索引不会越界
		if !node.IsEmpty() {
			fmt.Println(node.GetValue())
		}
	}

	list.PopFormTail(4)
	list.PopFormHead(2)

	fmt.Println("----------")

	for i := 0; i < list.GetLen(); i++ {
		// 从头部开始索引
		node := list.GetFormTail(i)

		// 节点为空不可能，因为list.Len()使得索引不会越界
		if !node.IsEmpty() {
			fmt.Println(node.GetValue())
		}
	}

	fmt.Println("----------")

}
