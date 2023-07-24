package ArrayAndString

import "strings"

func detectCapitalUse(word string) bool {
	//条件1：将所有字符转为大写
	Up := strings.ToUpper(word)
	Low := strings.ToLower(word)
	if Up == word {
		return true
	} else if Low == word { //条件2：将所有字符转为小写
		return true
	} else { //二者不满足的情况，看字符串是否只有一个大写字母且为首字母
		return HasUpper(word)
	}
}

func HasUpper(word string) bool {
	Up := "QWERTYUIOPASDFGHJKLZXCVBNM"
	tag := 0
	for i, w := range word {
		for j, u := range Up {
			if j == i && i == 0 && w == u {
				tag++
			} else if w == u {
				tag++
			}
		}
		if tag != 1 {
			return false
		}
	}
	return tag == 1
}
