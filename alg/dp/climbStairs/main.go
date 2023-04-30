package main

import "fmt"

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	//确定dp数组，dp[i]表示爬i阶楼梯的方法数
	dp := make([]int, n+1)

	//初始化dp数组，dp[1]，dp[2]表示分别爬一阶和二阶的方法数
	dp[1], dp[2] = 1, 2
	//遍历顺序：后续结果需要前序推导，所有从左至右遍历
	for i := 3; i <= n; i++ {
		//爬到i阶楼梯，只有两种：爬到i-1阶然后再爬1步，i-2同理
		dp[i] = dp[i-2] + dp[i-1]
	}
	return dp[n]
}

func main() {
	fmt.Println(climbStairs(2))
}
