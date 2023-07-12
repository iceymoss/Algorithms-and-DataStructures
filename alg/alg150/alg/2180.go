package main

func countEven(num int) int {
	total := 0
	for i := 1; i <= num; i++ {
		sum := 0
		//j /= 10 用于向高位推进，  sum += j % 10计算位数之和
		for j := i; j > 0; j /= 10 {
			sum += j % 10
		}
		if sum%2 == 0 {
			total++
		}
	}
	return total
}
