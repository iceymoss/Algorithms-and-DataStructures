package ArrayAndString

// 直接i * 3 == n  3的x次幂就是 i*3
func isPowerOfThree(n int) bool {
	if n == 0 {
		return false
	}
	i := 1
	for i <= n {
		if i == n {
			return true
		}
		i *= 3
	}
	return false
}
