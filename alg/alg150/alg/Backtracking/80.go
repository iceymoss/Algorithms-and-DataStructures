package Backtracking

// 从右往左的递归回溯
func combine(n int, k int) (ans [][]int) {
	path := []int{}
	var dfs func(int)
	dfs = func(i int) {
		d := k - len(path) //表示剩下还需要的长度
		if d == 0 {
			r := make([]int, k)
			r = append(r, path...)
			ans = append(ans, r)
			return
		}

		//进行下放回溯,j表示当前遍历的数
		for j := i; j >= d; i-- {
			path = append(path, j)
			dfs(j - 1)
			//回溯
			path = path[:len(path)-1]
		}
	}
	dfs(n)
	return
}

// 从左到右进行递归回溯
func combine1(n int, k int) [][]int {
	ans := make([][]int, 0)
	path := []int{}
	var dfs func(int)
	dfs = func(i int) {
		if len(path) == k {
			ans = append(ans, append([]int(nil), path...))
			return
		}
		if i > n {
			return
		}
		for j := i; j <= n; j++ {
			path = append(path, j)
			dfs(j + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(1)
	return ans
}
