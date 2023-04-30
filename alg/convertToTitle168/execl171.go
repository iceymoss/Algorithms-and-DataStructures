package main

// “FXSHRXW” 中的每个字母对应的序号分别是：[6,24,19,8,18,24,23]（其中A到Z分别对应1到26)则列名称对应的列序号为：
// 23×26^0+24×26^1+18×26^2+8×26^3+19×26^4+24×26^5+6×26^6=2147483647
// 本题的核心思想是进制转换，转换26进制
func titleToNumber(columnTitle string) int {
	res := 0
	mul := 1 //26^0
	for i := len(columnTitle) - 1; i >= 0; i-- {
		num := columnTitle[i] - 'A' + 1 //基数
		res += int(num) * mul
		mul *= 26
	}
	return res
}
