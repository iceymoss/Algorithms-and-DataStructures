package heap

import "sort"

type SmallestInfiniteSet struct {
	arr []int
}

func Constructor() SmallestInfiniteSet {
	arr := make([]int, 0, 1000)
	for i := 1; i <= 1000; i++ {
		arr = append(arr, i)
	}
	return SmallestInfiniteSet{
		arr: arr,
	}
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	cur := this.arr[0]
	this.arr = this.arr[1:]
	return cur
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	//先判断是否在优先队列中
	x := this.arr[0]
	if num < x {
		//如果小于对头就将该元素插入到对头
		this.arr = append([]int{num}, this.arr...)
	} else {
		//找到num的队列中的位置，SearchInts：如果num在arr中就返回num在位置，没有则返回一个插入的位置
		t := sort.SearchInts(this.arr, num)
		//这里需要考虑是否越界
		if t < len(this.arr) {
			if this.arr[t] != num { //不相等说明num不在arr中，我们需要插入
				//以t位置为分割点
				tailarr := append([]int{num}, this.arr[t:]...)
				this.arr = append(this.arr[:t], tailarr...)
			}
		} else {
			//越界情况，直接插入arr的尾巴
			this.arr = append(this.arr, num)
		}
	}
}

// 基于哈希表
type SmallestInfiniteSetMap map[int]bool

func ConstructorMap() SmallestInfiniteSetMap {
	return SmallestInfiniteSetMap{}
}

func (s SmallestInfiniteSetMap) PopSmallest() int {
	//从最小值开始，
	//例如：第一次调用：i=1, s[1]=false, !s[1]=true ==> s[2]=true, 直接返回i结束
	//第二次调用：i=1，直接i++, i=2 s[2]=false, !s[2]=true ==> s[2]=true, 直接返回i结束
	//以此类推即可
	for i := 1; ; i++ {
		if !s[i] {
			s[i] = true
			return i
		}
	}
}

func (s SmallestInfiniteSetMap) AddBack(n int) {
	//直接删除对应key-value就即可
	delete(s, n)
}
