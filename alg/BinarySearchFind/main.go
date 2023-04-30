package main

import "fmt"

func BinarySearchFind(arr []int, tage int) int {
	left, right := 0, len(arr)-1
	var res int
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] == arr[tage] {
			res = arr[tage]
			return res
		} else if arr[tage] > arr[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func Mysqrt(x int) int {
	r := x
	for r*r > x {
		r = (r + x/r) / 2
	}
	return r

}

func main() {
	//arr := make([]int, 200)
	//for i := 0; i < 200; i++ {
	//	arr[i] = i
	//}
	//fmt.Println(arr)
	//fmt.Println(BinarySearchFind(arr, 1))

	//fmt.Println(Sqrt(16))
	fmt.Println(Mysqrt(3))

}
