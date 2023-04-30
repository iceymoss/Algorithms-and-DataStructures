package main

import "fmt"

// leetcode746
// 注意题目描述：爬上台阶后才消费，楼等为i+1
func minCostClimbingStairs(cost []int) int {
	//dp数组：dp[i]表示到i阶最小消费
	dp := make([]int, len(cost)+1)

	//到i阶只有两种走法：
	//1、到i-1阶然后再走1阶，然后加上i-1阶向上走的消费
	//2、到i-2阶然后再走2阶，然后加上i-2阶向上走的消费
	//递推公式：dp[i] = min(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])

	//由递推公式可以知道：dp[2] = dp[1] + dp[0]; dp[3] = dp[2] + dp[1]
	//所以dp初识化
	dp[0], dp[1] = 0, 0

	//由递推公式可以知道必须从左向右遍历
	for i := 2; i <= len(cost); i++ {
		dp[i] = min(dp[i-2]+cost[i-2], dp[i-1]+cost[i-1])
	}
	fmt.Println(dp)
	return dp[len(dp)-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func main() {
	fmt.Println(minCostClimbingStairs([]int{10, 15, 20}))
}
