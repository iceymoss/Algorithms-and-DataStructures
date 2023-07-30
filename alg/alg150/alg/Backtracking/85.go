package Backtracking

func generateParenthesis(n int) []string {
	//回溯
	m := n * 2
	res := make([]string, 0)
	path := make([]byte, m)
	var dfs func(int, int)
	//i表示括号总数，open表示左括号总数
	dfs = func(i, open int) {
		if i == m {
			res = append(res, string(path))
			return
		}
		if open < n { //填左括号
			path[i] = '('
			dfs(i+1, open+1)
		}
		if i-open < open { //填右括号：i-open<open左括号数大于右括号数的情况可以填写右括号
			path[i] = ')'
			dfs(i+1, open)
		}
	}
	dfs(0, 0)
	return res
}
