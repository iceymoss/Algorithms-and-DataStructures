package DoublePointer

import (
	"sort"
)

//暴力枚举，使用哈希表进行标记

func MaxOperations(nums []int, k int) int {
	if len(nums) <= 0 {
		return 0
	}
	m := make(map[*int]bool)
	tag := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[j]+nums[i] == k && m[&nums[j]] == m[&nums[i]] && !m[&nums[j]] {
				m[&nums[j]], m[&nums[i]] = true, true
				tag++
			}
		}
	}
	return tag
}

// 哈希表

func MaxOperations1(nums []int, k int) int {
	if len(nums) <= 0 {
		return 0
	}
	tag := 0
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if m[k-nums[i]] > 0 {
			tag++
			m[k-nums[i]]--
		} else {
			m[nums[i]]++
		}
	}
	return tag
}

//双指针法，先对数组进行排序，然后如果nums[left]+nums[right]=k => left++, right--；
//如果nums[left]+nums[right]>k => right--
//如果nums[left]+nums[right]<k => left++

func MaxOperations2(nums []int, k int) int {
	if len(nums) <= 0 {
		return 0
	}
	//排序
	sort.Ints(nums)
	tag := 0
	left, right := 0, len(nums)-1
	for left < right {
		if nums[left]+nums[right] == k {
			tag++
			left++
			right--
		} else if nums[left]+nums[right] > k {
			right--
		} else {
			left++
		}
	}
	return tag
}
