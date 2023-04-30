package main

import "fmt"

//leetcode 63
//本题和62题一样，需要注意的是，当dp[i][j]有阻碍物时，需要将对应的dp[i][j] = 0,表示阻碍物，便于相邻坐标计算
//需要注意的还有初始化数组时，需要将第一列和第一行初识化的值初始化为1，但是需要注意：对应的[0,j]和[i,0]有阻碍物后，对应后续必须为零

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	row, col := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([][]int, len(obstacleGrid))
	for i := 0; i < row; i++ {
		dp[i] = make([]int, col)
	}
	//对应的[0,j]有阻碍物后，对应后续必须为零
	for j := 0; j < col; j++ {
		if obstacleGrid[0][j] == 1 {
			break
		}
		dp[0][j] = 1
	}
	//对应的[i,0]有阻碍物后，对应后续必须为零
	for i := 0; i < row; i++ {
		if obstacleGrid[i][0] == 1 {
			break
		}
		dp[i][0] = 1
	}
	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			if obstacleGrid[i][j] != 1 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[row-1][col-1]
}

func main() {
	arr := [][]int{{0, 0, 0, 1}, {0, 0, 0, 0}, {1, 0, 1, 0}, {0, 1, 1, 0}, {0, 0, 0, 0}}
	fmt.Println(uniquePathsWithObstacles(arr))
}
