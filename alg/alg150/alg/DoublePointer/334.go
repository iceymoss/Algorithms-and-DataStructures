package DoublePointer

// 反转字符
func reverseString(s []byte) {
	//双指针法
	l, r := 0, len(s)-1
	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
}
