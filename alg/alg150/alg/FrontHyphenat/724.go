package FrontHyphenat

import "fmt"

func PivotIndex(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}

	leftsum := 0
	rightsum := 0
	for i := 0; i < len(nums); i++ {
		leftsum += nums[i]
		rightsum = sum - leftsum + nums[i]
		fmt.Printf("i:%d, left:%d, right:%d\n", i, leftsum, rightsum)
		if rightsum == leftsum {
			return i
		}
	}
	return -1
}
