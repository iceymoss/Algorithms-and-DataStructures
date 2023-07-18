package main

// 枚举法，如果相邻的格子是海，则周长+1
func islandPerimeter(grid [][]int) int {
	c := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				c += isSea(i, j, grid)
			}
		}
	}
	return c
}

// 返回可用边
func isSea(x, y int, p [][]int) int {
	h := len(p)
	l := len(p[0])
	total := 0
	if x-1 < 0 || (x-1 >= 0 && p[x-1][y] == 0) {
		total++
	}
	if x+1 >= h || (x+1 < h && p[x+1][y] == 0) {
		total++
	}
	if y-1 < 0 || (y-1 >= 0 && p[x][y-1] == 0) {
		total++
	}
	if y+1 >= l || (y+1 < l && p[x][y+1] == 0) {
		total++
	}
	return total
}

//可以使用dp解决
