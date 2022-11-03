package doubleList

import (
	"sync"
)

//LinkNode 构造节点
type LinkNode struct {
	Next  *LinkNode
	Per   *LinkNode
	Value string
}

//DoubleList 构造列表
type DoubleList struct {
	Head *LinkNode
	Tail *LinkNode
	Len  int
	Lock sync.Mutex
}

//GetValue 获取节点值
func (node *LinkNode) GetValue() string {
	return node.Value
}

//GetPer 获取前驱节点
func (node *LinkNode) GetPer() *LinkNode {
	return node.Per
}

//GetNext 获取后驱节点
func (node *LinkNode) GetNext() *LinkNode {
	return node.Next
}

//IsEmptyPer 前驱节点为空
func (node *LinkNode) IsEmptyPer() bool {
	return node.Per == nil
}

//IsEmptyNext 后驱节点为空
func (node *LinkNode) IsEmptyNext() bool {
	return node.Next == nil
}

//IsEmpty 当前节点为空
func (node *LinkNode) IsEmpty() bool {
	return node == nil
}

func (list *DoubleList) GetLen() int {
	return list.Len
}

//AddNodeFromHead 从头部某个位置插入新节点
func (list *DoubleList) AddNodeFromHead(n int, element string) {
	//并发锁
	list.Lock.Lock()
	defer list.Lock.Unlock()

	//n=0,为列表头，依次类推
	cur := list.Head

	if n != 0 && n >= list.Len {
		panic("index out")
	}

	//向n位置插入节点，原n位置节点边为n+1
	//寻找需要插入的位置,插入到将原位置向后移
	for i := 1; i <= n; i++ {
		cur = cur.Next
	}

	newNode := new(LinkNode)
	newNode.Value = element

	//如果为空列表
	if cur.IsEmpty() {
		list.Head = newNode
		list.Tail = newNode
	} else {
		per := cur.Per
		//如果原节点的前驱为空，新节点为头节点
		if per.IsEmpty() {
			cur.Per = newNode
			newNode.Next = cur
			list.Head = newNode
		} else {
			per.Next = newNode
			newNode.Per = per

			newNode.Next = cur
			cur.Per = newNode
		}
	}
	list.Len++
}

//AddNodeFromTail 从尾部某个位置插入新节点
func (list *DoubleList) AddNodeFromTail(n int, element string) {
	list.Lock.Lock()
	defer list.Lock.Unlock()

	if n != 0 && n >= list.Len {
		panic("index out")
	}

	cur := list.Tail

	for i := 1; i <= n; i++ {
		cur = cur.Per
	}

	newNode := new(LinkNode)
	newNode.Value = element

	//列表为空
	if cur.IsEmpty() {
		list.Tail = newNode
		list.Head = newNode
	} else {

		next := cur.Next

		//只有一个节点的时候，新节点为列表尾部插入
		if next.IsEmpty() {
			cur.Next = newNode
			newNode.Per = cur
			list.Tail = newNode
		} else {
			next.Per = newNode
			newNode.Next = next

			cur.Next = newNode
			newNode.Per = cur
		}
	}
	list.Len++
}

//GetFormHead 从头部获取节点
func (list *DoubleList) GetFormHead(n int) *LinkNode {
	//不能越界
	if n < 0 || n >= list.Len {
		panic("index out")
	}

	cur := list.Head

	for i := 1; i <= n; i++ {
		cur = cur.Next
	}
	return cur
}

//GetFormTail 从尾部获取节点
func (list *DoubleList) GetFormTail(n int) *LinkNode {
	if n < 0 || n >= list.Len {
		panic("index out")
	}

	cur := list.Tail

	for i := 1; i <= n; i++ {
		cur = cur.Per
	}
	return cur
}

//PopFormHead 从头部移除某个位置的节点
func (list *DoubleList) PopFormHead(n int) *LinkNode {
	list.Lock.Lock()
	defer list.Lock.Unlock()

	if n < 0 || n >= list.Len {
		panic("index out")
	}

	cur := list.Head

	for i := 1; i <= n; i++ {
		cur = cur.Next
	}

	if cur.IsEmpty() {
		panic("list is empty")
	} else {
		next := cur.Next
		per := cur.Per

		if per.IsEmpty() {
			if list.Len == 1 {
				//只有一个节点
				list.Head = nil
				list.Tail = nil
			} else {
				//移除头节点，下一个节点变为头节点
				list.Head = next
				next.Per = nil
			}

		} else if next.IsEmpty() {
			//最后一个节点
			list.Tail = per
			cur.Per.Next = nil

		} else {
			//中间节点
			per.Next = next
			next.Per = per
		}
	}
	list.Len--
	return cur
}

//PopFormTail 从尾部移除某位置节点
func (list *DoubleList) PopFormTail(n int) *LinkNode {
	list.Lock.Lock()
	defer list.Lock.Unlock()

	if n < 0 || n >= list.Len {
		panic("index out")
	}

	cur := list.Tail

	for i := 1; i <= n; i++ {
		//移动当前节点
		cur = cur.Per
	}

	per := cur.Per
	next := cur.Next

	if per.IsEmpty() && next.IsEmpty() {
		//只有一个节点
		list.Head = nil
		list.Tail = nil
	} else if next.IsEmpty() {
		//移除尾部，原尾部的下一个节点，变成尾部
		list.Tail = per
		per.Next = nil
	} else if per.IsEmpty() {
		//移除头部
		list.Head = next
		next.Per = nil
	} else {
		//移除中间
		next.Per = per
		per.Next = next
	}
	list.Len--
	return cur
}
