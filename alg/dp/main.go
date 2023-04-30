package main

import "fmt"

func climbStairs(n int) int {
	tag := make([]int, n+1)
	return stairsdp(n, tag)
}

func stairsdp(n int, tag []int) int {
	if n == 2 {
		return 2
	}
	if n == 1 {
		return 1
	}

	if tag[n] != 0 {
		tag[n] = stairsdp(n-1, tag) + stairsdp(n-2, tag)
		return tag[n]
	}
	return stairsdp(n-1, tag) + stairsdp(n-2, tag)
}

func test(n int) int {
	//核心是统计爬楼梯方式的数量
	//用dp来记录到第i阶楼梯的方式数，使用dp记录
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i < n+1; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func test2(n int) int {
	cur := 1
	per := 1
	for i := 2; i <= n; i++ {
		per, cur = cur, per+cur
	}
	return cur
}

func main() {
	fmt.Println(test2(5))
}
