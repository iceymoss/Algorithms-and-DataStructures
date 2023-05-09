package HashesMap

//直接使用哈希表做记录即可

func FindDifference(nums1 []int, nums2 []int) [][]int {
	res := make([][]int, 2)
	for k, _ := range res {
		res[k] = make([]int, 0)
	}
	m1 := make(map[int]bool)
	for _, v := range nums2 {
		m1[v] = true
	}
	m2 := make(map[int]bool)
	for _, v := range nums1 {
		m2[v] = true
	}

	for _, v := range nums1 {
		if !m1[v] {
			res[0] = append(res[0], v)
			m1[v] = true
		}
	}

	for _, v := range nums2 {
		if !m2[v] {
			res[1] = append(res[1], v)
			m2[v] = true
		}
	}

	return res
}
