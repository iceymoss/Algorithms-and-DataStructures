package main

import (
	"fmt"
	"math"
)

func maxProfit(arr []int) int {
	//dp[i]表示前i日的最大利润
	//dp[i] = max(dp[i-1],p[i]-min[p[0:n])
	//dp[0] = 0
	//遍历顺序：有左到右
	dp := make([]int, len(arr))
	dp[0] = 0
	min := math.MaxInt
	for i := 1; i < len(arr); i++ {
		if min > arr[i-1] {
			min = arr[i-1]
		}
		dp[i] = max(dp[i-1], arr[i]-min)
	}
	return dp[len(arr)-1]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	fmt.Println(maxProfit([]int{1, 2}))

}
