package ArrayAndString

func isUgly(n int) bool {
	arr := []int{2, 3, 5}
	if n <= 0 {
		return false
	}
	for _, v := range arr {
		for n%v == 0 {
			n /= v
		}
	}
	return n == 1
}
