package ArrayAndString

func findWords(words []string) []string {
	//暴力枚举
	res := make([]string, 0)
	line1 := "qwertyuiopQWERTYUIOP"
	line2 := "asdfghjklASDFGHJKL"
	line3 := "zxcvbnmZXCVBNM"
	for _, word := range words {
		if IsCantant(word, line1) {
			res = append(res, word)
		}
		if IsCantant(word, line2) {
			res = append(res, word)
		}
		if IsCantant(word, line3) {
			res = append(res, word)
		}
	}
	return res
}

func IsCantant(str string, mod string) bool {
	tag := 0
	for _, s := range str {
		for _, l := range mod {
			if l == s {
				tag++
			}
		}
	}
	return len(str) == tag
}
