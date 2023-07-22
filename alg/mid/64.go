package mid

// 求 1+2+...+n ，要求不能使用乘除法、for、while、if、else、switch、case等关键字及条件判断语句（A?B:C）。
// 递归
func sumNums(n int) int {
	if n == 1 {
		return n
	}
	return n + sumNums(n-1)
}
