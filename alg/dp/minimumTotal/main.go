package main

import "fmt"

func minimumTotal(triangle [][]int) int {
	signIndex := 0
	deth := 0
	//遍历二维数组
	for i := 0; i < len(triangle); i++ {
		arr := triangle[i]
		if len(arr) == 1 {
			deth += arr[signIndex]
			continue
		}
		if arr[signIndex] <= arr[signIndex+1] {
			deth += arr[signIndex]
		} else {
			deth += arr[signIndex+1]
			signIndex = signIndex + 1
		}
	}
	return deth
}

func minimumTotal1(triangle [][]int) int {
	//回溯法(超时）无法剪枝
	return dfs(triangle, 0, 0)
}

func dfs(triangle [][]int, row int, signIndex int) int {
	//递归到底
	if row == len(triangle) {
		return 0
	}
	//下放-回溯
	ldeth := dfs(triangle, row+1, signIndex)
	rdeth := dfs(triangle, row+1, signIndex+1)

	if ldeth <= rdeth {
		return ldeth + triangle[row][signIndex]
	} else {
		return rdeth + triangle[row][signIndex]
	}
}

func minimumTotal2(triangle [][]int) int {
	//动态规划，自底向上，dp状态，dp方程
	h := len(triangle)
	dp := make([][]int, h)
	for i := 0; i < h; i++ {
		dp[i] = make([]int, len(triangle[i]))
	}

	//自底向上遍历二维数组
	for i := h - 1; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			//到达二维数组底层，将底层所有元素给到dp
			if i == h-1 {
				//dp状态
				dp[i][j] = triangle[i][j]
			} else {
				//dp状态方程
				//回到上一层时，需要判断的当前层的子层的最小值
				dp[i][j] = mini(dp[i+1][j], dp[i+1][j+1]) + triangle[i][j]
			}
		}
	}
	return dp[0][0]
}

func mini(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func main() {
	arr := [][]int{{-1}, {2, 3}, {1, -1, -5}}
	fmt.Println(minimumTotal2(arr))

}
