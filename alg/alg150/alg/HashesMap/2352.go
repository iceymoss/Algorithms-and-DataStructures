package HashesMap

import (
	"fmt"
	"strings"
)

func EqualPairs(grid [][]int) int {
	//使用map，先遍历行，再遍历列
	resTotal := 0
	row, col := len(grid), len(grid[0])
	M := make(map[string]int, 0)
	//遍历行
	for i := 0; i < row; i++ {
		str := []string{}
		for j := 0; j < col; j++ {
			//构造一个字符串
			str = append(str, fmt.Sprintf("%v", grid[i][j]))
		}
		ts := strings.Join(str, ",")
		fmt.Println("整合的字符串：", ts)
		M[ts]++
	}

	for k := 0; k < col; k++ {
		str := []string{}
		for v := 0; v < row; v++ {
			str = append(str, fmt.Sprintf("%v", grid[v][k]))
		}
		ts := strings.Join(str, ",")
		fmt.Println("字符串2：", ts)
		fmt.Println("结果：", resTotal)
		if M[ts] > 0 {
			resTotal += M[ts]
		}
	}
	return resTotal
}
