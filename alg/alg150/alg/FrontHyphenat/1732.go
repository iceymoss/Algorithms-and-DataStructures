package FrontHyphenat

import "fmt"

//题目给定的是变化量，我们维护一个和即可

func LargestAltitude(gain []int) int {
	sum, mtin := 0, 0
	for i := 0; i < len(gain); i++ {
		sum += gain[i]
		fmt.Println(sum)
		if sum > mtin {
			mtin = sum
		}
	}
	return mtin
}
