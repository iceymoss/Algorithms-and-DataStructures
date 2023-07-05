package ArrayAndString

import "fmt"

func summaryRanges(nums []int) []string {
	result := make([]string, 0)
	if len(nums) == 0 {
		return result
	}
	start := nums[0]
	v := nums[0]
	for i := 1; i < len(nums); i++ {
		if (v + 1) == nums[i] {
			v = nums[i]
		} else {
			if start == v {
				result = append(result, fmt.Sprintf("%d", start))
			} else {
				result = append(result, fmt.Sprintf("%d->%d", start, nums[i-1]))
			}
			start = nums[i]
			v = nums[i]
		}
	}
	if start == v {
		result = append(result, fmt.Sprintf("%d", v))
	} else {
		result = append(result, fmt.Sprintf("%d->%d", start, v))
	}
	return result
}
