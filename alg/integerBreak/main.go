package main

import "fmt"

func integerBreak(n int) int {
	//dp[i]表示i拆分后的k个数之积
	dp := make([]int, n+1)

	//要想得到dp[i],想要计算j*(i-j)和j*dp[i-j]的值，j表示从1~i-1
	//对于每一个dp[i]都要计算i和的k个子数，从而找出最大值
	//状态方程：dp[i] = max(dp[i], max(j*(i-j), j*dp[i-j]))

	//初始化：dp[0],d[1]是无意义的，初始化dp[2] = 1
	dp[2] = 1
	//遍历顺序：dp[i],需要靠dp[i-j]推导，从左向右遍历
	for i := 3; i <= n; i++ {
		for j := 1; j < n-1; j++ {
			dp[i] = max(dp[i], max(j*(i-j), j*dp[i-j]))
		}
	}
	return dp[n]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	fmt.Println(integerBreak(3))

}
