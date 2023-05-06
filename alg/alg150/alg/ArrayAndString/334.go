package ArrayAndString

import "math"

//暴力法:直接枚举

func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i] < nums[j] && nums[j] < nums[k] {
					return true
				}
			}
		}
	}
	return false
}

//可以使用贪心的方法将空间复杂度降到 O(1)。从左到右遍历数组 nums，遍历过程中维护两个变量 first和 second，
//分别表示递增的三元子序列中的第一个数和第二个数，任何时候都有 first<second。
//初始时，first=nums[0],second=+∞对于 1≤i<n1，当遍历到下标 i 时，令 num=nums[i]进行如下操作：

//赋初始值的时候，已经满足second > first了，现在找第三个数third
//(1) 如果third比second大，那就是找到了，直接返回true
//(2) 如果third比second小，但是比first大，那就把second指向third，然后继续遍历找third
//(3) 如果third比first还小，那就把first指向third，然后继续遍历找third（这样的话first会跑到second的后边，但是不要紧，因为在second的前边，老first还是满足的）

func increasingTriplet1(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	first, second := nums[0], math.MaxInt
	for i := 1; i < len(nums); i++ {
		num := nums[i]
		if num > second {
			return true
		} else if num > first {
			second = num
		} else {
			first = num
		}
	}

	return false
}

//DP:我们遍历数组0——n-1，遍历到i时我们需要在i的左边找到一个最小值，在i右边找到一个最大值，最后满足min[i-1] < nums[i] < max[i+]
//所以需要使用动态规划来找出i左边的最小值，i右边的最大值，动态规划五部曲：
//1、dp[i]的含义：
//	这里我们要找最小值和最大值：
//		dpMin[i]表示从0-i中nums数组的最小值，例如:nums = [10,200,2,10] => dpMin = [10,10,2,2]
//		dpMax[i]表示从len(nums)-i 的最大值，例如:nums= = [10,5,13] => dpMax = [13,13,13]

//2、状态转移方程：
//		dpMin[i] = min(dpMin[i-1],nums[i])
//		dpMax[i] = max(dpMax[i+1],nums[i])

//3、初始化dp数组：dpMin[0] = nums[0], dpMax[len(nums)-1] = nums[len(nums)-1]

//4、遍历顺序：由方程和初始化，看出：dpMin：L -> R;   dpMax：R -> L

//5、根据情况打印dp数组

func increasingTripletDP(nums []int) bool {
	n := len(nums)
	dpMin, dpMax := make([]int, n), make([]int, n)
	dpMin[0], dpMax[n-1] = nums[0], nums[n-1]

	for i := 1; i < n; i++ {
		dpMin[i] = min(dpMin[i-1], nums[i])
	}

	for j := n - 2; j >= 0; j-- {
		dpMax[j] = max(dpMax[j+1], nums[j])
	}

	for k := 1; k < n-1; k++ {
		if dpMin[k-1] < nums[k] && nums[k] < dpMax[k+1] {
			return true
		}
	}
	return false
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
