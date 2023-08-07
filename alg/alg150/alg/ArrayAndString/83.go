package ArrayAndString

import "strconv"

func CountAndSay(n int) string {
	ans := "1"
	for i := 1; i < n; i++ {
		per := ""
		for j, start := 0, 0; j < len(ans); start = j {
			//统计相同值的数量
			for j < len(ans) && string(ans[start]) == string(ans[j]) {
				j++
			}
			per += strconv.Itoa(j-start) + string(ans[start])
			//将tag和拼接
		}
		ans = per
	}
	return ans
}
