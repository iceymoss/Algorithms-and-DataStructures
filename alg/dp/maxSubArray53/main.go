package main

import (
	"fmt"
	"math"
)

// 暴力法：从最大数量子数组开始(从本身开始）在该子数组中相邻的找到那几个元素之和最大
// 使用两层for，第一层for完成数组在收缩，第二层for计算当前子数组和最大值
func maxSubArray(nums []int) int {
	res := math.MinInt
	for i := 0; i < len(nums); i++ {
		count := 0
		for j := i; j < len(nums); j++ {
			count += nums[j]
			if count > res {
				res = count
			}
		}
	}
	return res
}

// 贪心法：局部最优到全局最优，遍历数组，局部最优：当前“连续和”为负数的时候立刻放弃，
// 从下一个元素重新计算“连续和”，因为负数加上下一个元素 “连续和”只会越来越小。
// 例如：x为负数，y为正数， 一定有 x+y < y
func maxSubArray1(nums []int) int {
	res := math.MinInt               //返回值最小
	count := 0                       //用于计数
	for i := 0; i < len(nums); i++ { //遍历数组
		count += nums[i] //当遍历到元素相加
		if res < count { //找到最大值
			res = count
		}
		if count <= 0 { //当连续小于0时
			count = 0
		}
	}
	return res
}

// 动态规划：
func maxSubArray2(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = nums[0]
	res := nums[0] //dp状态
	for i := 1; i < n; i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i]) //dp方程
		res = max(res, dp[i])
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray2(arr))
}
