package heap

type heap struct {
	arr   []int
	count int
}

func new() *heap {
	//变量值为1
	arr := make([]int, 1)
	return &heap{
		arr:   arr,
		count: 0,
	}
}

func (h *heap) insert(key int) {
	h.arr = append(h.arr, key)
	h.count++
	h.shitUP(h.count)
}

func (h *heap) revmoe() int {
	if h.count < 1 {
		return -1
	}
	max := h.arr[1]
	h.arr[1], h.arr[h.count] = h.arr[h.count], h.arr[1]
	h.arr = h.arr[:len(h.arr)-1]
	h.count--
	h.shitDown(1)
	return max
}

func (h *heap) shitUP(index int) {
	for index > 1 && h.arr[index/2] < h.arr[index] {
		h.arr[index/2], h.arr[index] = h.arr[index], h.arr[index/2]
		index /= 2
	}
}

func (h *heap) shitDown(index int) {
	for 2*index < h.count {
		j := 2 * index
		if j+1 < h.count && h.arr[j+1] > h.arr[j] {
			j++
		}
		if h.arr[index] > h.arr[j] {
			break
		}
		h.arr[index], h.arr[j] = h.arr[j], h.arr[index]
		index = j
	}
}

func FindKthLargest(nums []int, k int) int {
	//堆排序
	maxheap := new()
	for _, v := range nums {
		maxheap.insert(v)
	}

	for i := 0; maxheap.count != 0 && i < len(nums); i++ {
		nums[i] = maxheap.revmoe()
	}

	return nums[k]
}
