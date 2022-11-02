package slice

import (
	"sync"
)

//Array 自定义一个可变长数组
type Array struct {
	array []int
	len   int
	cap   int
	lock  sync.Mutex
}

//Make 新建一个可变长数组
func Make(len, cap int) *Array {
	s := new(Array)
	if len > cap {
		panic("len large than cap")
	}

	s.len = 0
	s.cap = cap

	//将slice当中定长数组使用
	arr := make([]int, cap, cap)

	s.array = arr
	return s
}

func (a *Array) Append(element int) {
	a.lock.Lock()
	defer a.lock.Unlock()

	//如果容量不够，则扩容
	if a.len == a.cap {
		NewCap := a.len * 2

		//没有容量这加一个容量
		if a.cap == 0 {
			NewCap = 1
		}

		arr := make([]int, NewCap, NewCap)

		//将原数组放入新数组
		for k, v := range a.array {
			arr[k] = v
		}

		a.array = arr
		a.cap = NewCap
	}

	a.array[a.len] = element
	a.len++
}

//AppendMany 添加多个数据
func (a *Array) AppendMany(element ...int) {
	for _, v := range element {
		a.Append(v)
	}
}

func (a *Array) Get(index int) int {
	if index < 0 || index > a.len {
		panic("index large than len")
	}
	return a.array[index]
}

func (a *Array) Len() int {
	return a.len
}

func (a *Array) Cap() int {
	return a.cap
}
