package mid

// 回溯法
func permute(nums []int) [][]int {
	n := len(nums)
	an := make([]int, n)
	ans := make([][]int, 0)

	//记录是数组对应位置否被标记
	onPath := make([]bool, n)
	var dfs func(int)

	dfs = func(i int) {
		if i == n {
			ans = append(ans, an)
			return
		}
		for j, on := range onPath {
			if !on {
				an[i] = nums[j]
				onPath[j] = true
				dfs(i + 1)
				onPath[j] = false
			}
		}
	}
	dfs(0)
	return ans
}
