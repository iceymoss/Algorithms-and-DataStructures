package ArrayAndString

func wordPattern(pattern string, s string) bool {
	//处理这个s
	str := make([]string, 0)
	st := ""
	for i := 0; i < len(s); i++ {
		if string(s[i]) == " " {
			str = append(str, st)
			st = ""
		} else if i == len(s)-1 {
			st += string(s[i])
			str = append(str, st)
			break
		} else {
			st += string(s[i])
		}
	}

	if len(pattern) != len(str) {
		return false
	}

	//使用哈希做映射
	pMap, sMap := make(map[byte]string), make(map[string]byte)
	for j := 0; j < len(str); j++ {
		if _, ok := pMap[pattern[j]]; !ok {
			pMap[pattern[j]] = str[j]
		}
		if _, ok := sMap[str[j]]; !ok {
			sMap[str[j]] = pattern[j]
		}

		if pMap[pattern[j]] != str[j] || sMap[str[j]] != pattern[j] {
			return false
		}
	}
	return true
}
