package HashesMap

import "sort"

func closeStrings(word1 string, word2 string) bool {

	//根据题意，我们需要进行以下判断：
	//1、两个字符串必须相等
	//2、两个字符串中的字符类型必须相同
	//      满足操作一
	//3、两个字符串字符串出现的频次必须相同
	//      满足操作一和操作二

	//长度必须相同
	if len(word1) != len(word2) {
		return false
	}

	//两个字符串出现的字符类型必须相同
	//使用f1，f2分别记录word1和word2的26个字母出现的次数
	f1, f2 := make([]int, 26), make([]int, 26)
	for k, _ := range word1 {
		f1[word1[k]-'a']++
		f2[word2[k]-'a']++
	}

	for k, _ := range f1 {
		//判断两个数组对应位置否为有字符串元素
		if f1[k] == f2[k] || f1[k] > 0 && f2[k] > 0 {
			continue
		} else {
			return false
		}
	}

	//排序
	//判断字符串频次是否相等
	sort.Ints(f1)
	sort.Ints(f2)
	for k, _ := range f1 {
		if f1[k] == f2[k] {
			continue
		} else {
			return false
		}
	}
	return true
}
