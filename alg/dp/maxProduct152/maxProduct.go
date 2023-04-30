package main

import (
	"fmt"
	"math"
)

// 暴力法: 从最大数量子数组开始(从本身开始）在该子数组中相邻的找到那几个元素乘积最大值
// 使用两层for，第一层for完成数组的收缩，第二层for计算当前子数组乘机最大值
func maxProduct(nums []int) int {
	res := math.MinInt
	for i := 0; i < len(nums); i++ {
		count := 1
		for j := i; j < len(nums); j++ {
			count *= nums[j]
			if count > res {
				res = count
			}
		}
	}
	return res
}

// 贪心法
func maxProduct1(nums []int) int {
	res := math.MinInt
	count := 1
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			if res < nums[i] {
				res = nums[i]
			}
			count = 1
			continue
		}
		count *= nums[i]
		if res < count {
			res = count
		}
	}
	return res
}

// 回溯法:归并遍历所有可能的乘积， 然后将最大值回溯，但是会超时
func maxProduct2(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	res := nums[0]

	var dfs func(i int, val int)
	dfs = func(i int, val int) {
		if i >= len(nums)-1 {
			res = max(res, val)
			return
		}
		res = max(res, val)
		dfs(i+1, val*nums[i+1])
		dfs(i+1, nums[i+1])
	}
	dfs(0, nums[0])
	return res
}

// 动态规划:由于整数数组 nums 中的元素可能有正数、负数和 0，因此连续子数组中的元素也可能是这三种情况。
// 如果连续子数组中的元素存在负数，正数乘以负数就成负数，
// 那么最大值乘以负数就变成了最小值，因此需要同时考虑当前连续子数组乘积的最大值 curMax 和最小值 curMin。
// 注意点 整数数组 nums 中存在负数，当遍历到以 nums[i]（负数）结尾连续子数组时，需要交换 curMax 和 curMin。
func maxProduct3(nums []int) int {
	perMax, perMin, rMax := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			//如果出现负数，需要将最大最小值互换
			perMax, perMin = perMin, perMax
		}
		//记录最大乘机
		perMax = max(perMax*nums[i], nums[i])
		//记录最小乘机
		perMin = min(perMin*nums[i], nums[i])
		//获取最大乘机
		rMax = max(perMax, perMin)
	}
	return rMax
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	arr := []int{2, 0, -4}
	fmt.Println(maxProduct2(arr))
}
