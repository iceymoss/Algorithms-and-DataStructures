package SinglyLinkedLists

import "fmt"

//Node 节点
type Node struct {
	Data     int
	NextNode *Node
}

//LinkList 链表结构
type LinkList struct {
	Head   *Node
	Length int
}

//NewLinkList 初始化空链表
func NewLinkList() *LinkList {
	return &LinkList{}
}

//GetHead 获取head
func (l *LinkList) GetHead() int {
	return l.Head.Data
}

//IsEmpty 链表是否为空
func (l LinkList) IsEmpty() bool {
	if l.Head == nil {
		return true
	}
	return false
}

//Lengths 获取链表长度
func (l *LinkList) Lengths() int {
	count := 0
	current := l.Head
	for {
		if current.NextNode == nil {
			break
		}
		count++
		current = current.NextNode
	}
	count++
	return count
}

//InsertHead 向链表头部插入节点
func (l *LinkList) InsertHead(data int) {
	node := &Node{Data: data}
	if l.IsEmpty() {
		l.Head = node

	} else {
		node.NextNode = l.Head
		l.Head = node
		l.Length++
	}
	l.Length++
}

//InsertTail 向链表尾部插入节点
func (l *LinkList) InsertTail(data int) {
	node := &Node{Data: data}
	if l.Head == nil {
		l.Head = node
	} else {
		current := l.Head
		for {
			if current.NextNode == nil {
				current.NextNode = node
				break
			}
			current = current.NextNode
		}
	}
	l.Length++
}

//Traverse 遍历链表
func (l *LinkList) Traverse() {
	current := l.Head
	for current != nil {
		if current.NextNode == nil {
			fmt.Println(current.Data)
			return
		}
		fmt.Println(current.Data)
		current = current.NextNode
	}
}

//Insert 向指定位置插入节点
func (l *LinkList) Insert(i, data int) error {
	if i < 0 || i > l.Length {
		return fmt.Errorf("位置越界")
	}
	node := Node{Data: data, NextNode: nil}
	current := l.Head
	index := 0
	if i == 0 {
		node.NextNode = l.Head
		l.Head = &node
		l.Length++
		return nil
	}

	for index < i-1 {
		index++
		current = current.NextNode
	}
	tailNode := current.NextNode
	current.NextNode = &node
	node.NextNode = tailNode

	l.Length++
	return nil
}

//Remove 移除指定节点
func (l *LinkList) Remove(i int) int {
	if i < 0 || i > l.Length {
		return -1
	}
	current := l.Head
	index := 0

	if i == 0 {
		l.Head = current.NextNode
		l.Length--
		return current.Data
	}

	for index < i-1 {
		current = current.NextNode
		index++
	}
	data := current.NextNode.Data
	current.NextNode = current.NextNode.NextNode
	l.Length--
	return data
}

//GetIndex 获取指定节点索引
func (l *LinkList) GetIndex(n int) int {
	current := l.Head
	index := 0
	for index < l.Length-1 {
		if current.Data == n {
			return index
		}
		index++
		current = current.NextNode
	}
	return -1
}
