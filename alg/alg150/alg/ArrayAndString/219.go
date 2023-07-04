package ArrayAndString

// 暴力枚举
func containsNearbyDuplicate(nums []int, k int) bool {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] && abs(j, i) <= k {
				return true
			}
		}
	}
	return false
}

// map快速判复，然后判断abs(i-j)<=k，有则直接返回，没有则更新值
func containsNearbyDuplicateToMap(nums []int, k int) bool {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := m[nums[i]]; !ok {
			m[nums[i]] = i
		} else {
			if abs(i, m[nums[i]]) <= k {
				return true
			}
			m[nums[i]] = i
		}
	}
	return false
}

func abs(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}
