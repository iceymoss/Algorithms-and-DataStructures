package Stack

func removeStars(s string) string {
	//使用栈
	if len(s) == 0 {
		return ""
	}
	stack := make([]byte, 0)
	for _, v := range s {
		if v == '*' && len(stack) > 0 {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, byte(v))
		}
	}
	return string(stack)
}
