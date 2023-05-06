package SlidingWindow

import "strings"

//滑动窗口，使用滑动窗口

func MaxVowels(s string, k int) int {
	pSum := 0
	str := "aeiou"
	//窗口从左到右滑动
	//先设定初始值，计算窗口中的元音字母个数
	for i, _ := range s[:k] {
		if strings.Contains(str, string(s[i])) {
			pSum++
		}
	}
	resMax := pSum
	//滑动窗口，每次想窗口中加入一个元素和移除一个元素
	for i := k; i < len(s); i++ {
		//判断移除元素是否为元音
		if strings.Contains(str, string(s[i-k])) {
			pSum--
		}
		//判断加入元素是否为元音
		if strings.Contains(str, string(s[i])) {
			pSum++
		}
		resMax = max(resMax, pSum)
	}
	return resMax
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func MaxVowels1(s string, k int) int {
	pSum := 0
	//窗口从左到右滑动
	//先设定初始值，计算窗口中的元音字母个数
	for i, _ := range s[:k] {
		if isHas(s[i]) {
			pSum++
		}
	}

	resMax := pSum
	//滑动窗口，每次想窗口中加入一个元素和移除一个元素
	for i := k; i < len(s); i++ {
		//判断移除元素是否为元音
		if isHas(s[i-k]) {
			pSum--
		}
		//判断加入元素是否为元音
		if isHas(s[i]) {
			pSum++
		}
		resMax = max(resMax, pSum)
	}
	return resMax
}

func isHas(s byte) bool {
	if s == 'a' || s == 'e' || s == 'i' || s == 'o' || s == 'u' {
		return true
	}
	return false
}
