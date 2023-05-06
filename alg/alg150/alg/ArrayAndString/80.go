package ArrayAndString

// 快慢指针法
func removeDuplicates(nums []int) int {
	//简单的去重
	if len(nums) < 3 {
		return len(nums)
	}
	//快慢指针
	fast, slow := 2, 2
	for fast < len(nums) {
		if nums[slow-2] != nums[fast] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}

func remver(nums []int) int {
	if len(nums) < 3 {
		return len(nums)
	}
	var perocce func(k int) int
	perocce = func(k int) int {
		u := 0
		for _, v := range nums {
			if u < k || nums[u-k] != v {
				nums[u] = v
			}
			u++
		}
		return u
	}
	return perocce(2)
}
