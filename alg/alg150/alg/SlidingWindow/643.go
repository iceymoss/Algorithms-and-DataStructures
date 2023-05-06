package SlidingWindow

import "fmt"

//滑动窗口

func findMaxAverage(nums []int, k int) float64 {
	sum := 0
	for _, v := range nums[:k] {
		sum += v
		fmt.Println(v)
	}
	maxV := float64(sum)
	for i := k; i < len(nums); i++ {
		sum = sum - nums[i-k] + nums[i]
		maxV = max1(maxV, float64(sum))
	}
	return maxV / float64(k)
}

func max1(x, y float64) float64 {
	if x > y {
		return x
	}
	return y
}
