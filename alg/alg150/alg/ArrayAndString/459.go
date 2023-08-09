package ArrayAndString

func repeatedSubstringPattern(s string) bool {
	n := len(s)
	//第一个for获取到子串
	for i := 1; i*2 <= n; i++ {
		if n%i == 0 {
			match := true
			//第二个for判断s是否由子串构成
			for j := i; j < n; j++ {
				if s[j] != s[j-i] {
					match = false
					break
				}
			}
			if match {
				return true
			}
		}
	}
	return false
}
