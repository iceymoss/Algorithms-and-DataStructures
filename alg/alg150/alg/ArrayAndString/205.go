package ArrayAndString

// 思路：两字符之间相互映射：例：edd -> add, ado -> aff
func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sm, tm := make(map[byte]byte), make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		if _, ok := sm[s[i]]; !ok {
			sm[s[i]] = t[i]
		}

		if _, ok := tm[t[i]]; !ok {
			tm[t[i]] = s[i]
		}

		if sm[s[i]] != t[i] || tm[t[i]] != s[i] {
			return false
		}
	}
	return true
}
