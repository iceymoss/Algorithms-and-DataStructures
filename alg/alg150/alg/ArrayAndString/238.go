package ArrayAndString

func ProductExceptSelf(nums []int) []int {
	ans := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		ans[i] = 1
		for j := 0; j < len(nums); j++ {
			if i == j {
				continue
			}
			ans[i] *= nums[j]
		}
	}
	return ans
}

func ProductExceptSelf1(nums []int) []int {
	ans := make([]int, len(nums))
	total := 1
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			continue
		}
		total *= nums[i]
	}
	for j := 0; j < len(nums); j++ {
		if nums[j] == 0 {
			ans[j] = 0
			continue
		}
		ans[j] = total / nums[j]
	}
	return ans
}

func ProductExceptSelf2(nums []int) []int {
	n := len(nums)
	L, R, answer := make([]int, n), make([]int, n), make([]int, n)
	//0的左边和n-1的右边都没有元素所以都为1
	L[0], R[n-1] = 1, 1

	for i := 1; i < n; i++ {
		L[i] = L[i-1] * nums[i-1]
	}

	for j := n - 2; j >= 0; j-- {
		R[j] = R[j+1] * nums[j+1]
	}

	for k := 0; k < n; k++ {
		answer[k] = L[k] * R[k]
	}
	return answer
}
