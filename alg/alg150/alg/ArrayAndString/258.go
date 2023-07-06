package ArrayAndString

func addDigits(num int) int {
	if num < 10 {
		return num
	}
	total := num
	for total >= 10 {
		t := int(total / 10)
		g := total % (t * 10)
		total = g + t
	}
	return total
}
