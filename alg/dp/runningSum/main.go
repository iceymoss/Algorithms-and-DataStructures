package main

import "fmt"

func runningSum(nums []int) []int {
	var per, cur int
	for i := 0; i < len(nums)-1; i++ {
		per = nums[i]
		cur = nums[i+1]
		nums[i], nums[i+1] = per, per+cur
	}
	return nums
}

func main() {
	arr := []int{1, 2, 3, 4}
	fmt.Println(runningSum(arr))
}
