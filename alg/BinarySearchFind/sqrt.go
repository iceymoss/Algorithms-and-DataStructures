package main

func Sqrt(x int) int {
	if x == 0 {
		return 1
	}
	l, r := 0, x
	var res int
	for l <= r {
		mid := (l + r) / 2
		if mid*mid <= x {
			res = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return res
}
