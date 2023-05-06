package DoublePointer

//双指针法，使用tp指针指向s,使用指针i指向t, 遍历t,比较s[tp] == t[i]

func IsSubsequence(s string, t string) bool {
	if (len(t) <= 0 && len(s) <= 0) || (len(t) > 0 && len(s) <= 0) {
		return true
	}
	tp := 0
	for i := 0; i < len(t); i++ {
		if s[tp] == t[i] {
			tp++
			if tp == len(s) {
				return true
			}
		}
	}
	return false
}
