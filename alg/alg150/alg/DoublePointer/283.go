package DoublePointer

import (
	"fmt"
)

func MoveZeroes(nums []int) {
	if len(nums) <= 1 {
		return
	}

	lows, fast := 0, 0
	for fast < len(nums) {
		if nums[lows] == 0 && nums[fast] != 0 {
			nums[lows], nums[fast] = nums[fast], nums[lows]
			fast++
			lows++
		} else if nums[lows] != 0 && nums[fast] != 0 {
			fast++
			lows++
		} else if nums[fast] == 0 {
			fast++
		}

	}
	fmt.Println(nums)
}

func MoveZeroes1(nums []int) {
	if len(nums) <= 1 {
		return
	}
	queue := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			queue = append(queue, nums[i])
		}
	}

	for j := 0; j < len(nums); j++ {
		if len(queue) != 0 {
			nums[j] = queue[0]
			queue = queue[1:]
		} else {
			nums[j] = 0
		}

	}
}
