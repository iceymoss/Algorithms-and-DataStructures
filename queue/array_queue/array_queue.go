package array_queue

import "sync"

type ArrayQueue struct {
	Array []int
	Size  int
	Lock  sync.Mutex
}

//Push 入队
func (q *ArrayQueue) Push(v int) {
	q.Lock.Lock()
	defer q.Lock.Unlock()

	q.Array = append(q.Array, v)
	q.Size++
}

func (q *ArrayQueue) Pop() int {
	q.Lock.Lock()
	defer q.Lock.Unlock()

	if q.Size == 0 {
		panic("empty")
	}

	/*
		//直接原位移动，但不会释放内存
		v := q.Array[0]
		for i := 1; i < q.Size; i++ {
			q.Array[i-1] = q.Array[i]
		}

		//原数组缩容
		q.Array = q.Array[0: q.Size-1]
	*/
	v := q.Array[0]
	newArray := make([]int, q.Size-1, q.Size-1)
	for i := 0; i < q.Size-1; i++ {
		newArray[i] = q.Array[i+1]
	}

	q.Array = newArray
	q.Size--

	return v
}

func (q *ArrayQueue) Peek() int {
	return q.Array[0]
}

func (q *ArrayQueue) IsEmpty() bool {
	return q.Size == 0
}

func (q *ArrayQueue) Len() int {
	return q.Size
}
