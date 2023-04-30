package main

import "fmt"

// 递归法
func fib(n int) int {
	if n <= 0 {
		return n
	}
	return fib(n-2) + fib(n-1)
}

// 迭代法
func fib1(n int) int {
	per := 0
	cur := 1
	for i := 0; i < n; i++ {
		per, cur = cur, per+cur
	}
	return per
}

// 动态规划
func fib2(n int) int {
	//确定dp数组及下标含义
	//dp递推公式:dp[i] = dp[i-2]+dp[i-1]
	//dp数组初始化
	//遍历数组：0-n遍历
	//打印dp数组
	if n <= 1 {
		return n
	}
	dp := make([]int, n+1)
	dp[0], dp[1] = 0, 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-2] + dp[i-1]
	}
	fmt.Println(dp)
	return dp[n]
}

func main() {
	fmt.Println(fib(5))
}
