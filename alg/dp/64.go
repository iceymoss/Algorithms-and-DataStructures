package main

// dp五部法则：
// 1、dp[i][j]表示从(0,0)->(i, j)的最小路径
// 2、状态转移方程：dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
// 3、初始化：要想获取dp[i][j]的值，必须知道dp[i][0]和dp[0][j] 即初始化初始位置的行和列
// 4、遍历方向由左到右、有上到下
// 5、打印值验证
func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	rows, columns := len(grid), len(grid[0])
	dp := make([][]int, len(grid))
	//初始化dp数组
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, columns)
	}

	//初始化
	dp[0][0] = grid[0][0]
	for i := 1; i < rows; i++ {
		dp[i][0] = dp[i-1][0] + grid[i][0]
	}
	for j := 1; j < columns; j++ {
		dp[0][j] = dp[0][j-1] + grid[0][j]
	}

	//遍历目标数组，按照行进行计算
	for i := 1; i < rows; i++ {
		for j := 1; j < columns; j++ {
			dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}
	return dp[rows-1][columns-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
