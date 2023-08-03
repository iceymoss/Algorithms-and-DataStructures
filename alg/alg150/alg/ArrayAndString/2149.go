package ArrayAndString

// func rearrangeArray(nums []int) []int {
//     nums1 := make([]int, 0)
//     nums2 := make([]int, 0)
//     for _, v := range nums {
//         if v > 0 {
//             nums1 = append(nums1, v)
//         }else {
//             nums2 = append(nums2, v)
//         }
//     }
//     j := 0
//     for i := 0; i < len(nums1); i++ {
//         nums[j] = nums1[i]
//         j++
//         nums[j] = nums2[i]
//         j++
//     }
//     return nums
// }

func rearrangeArray(nums []int) []int {
	ans := []int{}
	pos, neg := 0, 0
	for i := 0; i+i < len(nums); i++ {
		//找到正数
		for ; nums[pos] < 0; pos++ {
		}
		ans = append(ans, nums[pos])
		pos++
		for ; nums[neg] > 0; neg++ {
		}
		ans = append(ans, nums[neg])
		neg++
	}
	return ans
}
