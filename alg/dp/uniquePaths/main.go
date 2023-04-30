package main

import "fmt"

func uniquePaths(m int, n int) int {
	//dp[m][n]表示到坐标为【m,n】的路线总数
	//dp[m][n] = dp[m-1][n] + dp[m][n-1]
	//dp[0][n]=1, dp[m][0]=1
	//由递推公式得到，从左上往右下遍历
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		for j := 0; j < n; j++ {
			dp[i][0] = 1
			dp[0][j] = 1
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}

func main() {
	fmt.Println(uniquePaths(3, 2))
}
