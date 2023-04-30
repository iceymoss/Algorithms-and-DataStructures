package main

import "fmt"

// 本题的核心思想其实是进制转换，26进制的转换，但是比较特殊的是范围[1,26]，不是正常的的26进制的[0,26)
func convertToTitle(columnNumber int) string {
	res := make([]byte, 0)
	for columnNumber > 0 {
		//范围[1,26]需要偏移1数值
		columnNumber--
		//需要将取余的得到的基数转换位大写字母
		res = append(res, byte(columnNumber%26)+'A')
		columnNumber /= 26
	}

	//反转
	resStr := make([]byte, 0)
	for i := len(res) - 1; i >= 0; i-- {
		resStr = append(resStr, res[i])
	}
	return string(resStr)
}

func main() {
	fmt.Println(titleToNumber("ZY"))
}
