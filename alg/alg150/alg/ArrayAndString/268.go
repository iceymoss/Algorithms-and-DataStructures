package ArrayAndString

func missingNumber(nums []int) int {
	for i := 0; i <= len(nums); i++ {
		if !contain(i, nums) {
			return i
		}
	}
	return 0
}

func contain(x int, nums []int) bool {
	for _, v := range nums {
		if x == v {
			return true
		}
	}
	return false
}
