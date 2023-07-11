package ArrayAndString

func intersection(nums1 []int, nums2 []int) []int {
	res := make(map[int]int)
	for _, v := range nums1 {
		if contan(v, nums2) {
			res[v] = v
		}
	}
	arr := make([]int, 0)
	for _, v := range res {
		arr = append(arr, v)
	}
	return arr
}

func contan(tag int, nums []int) bool {
	for _, v := range nums {
		if tag == v {
			return true
		}
	}
	return false
}

func intersectionMap(nums1 []int, nums2 []int) (intersection []int) {
	set1 := map[int]struct{}{}
	for _, v := range nums1 {
		set1[v] = struct{}{}
	}
	set2 := map[int]struct{}{}
	for _, v := range nums2 {
		set2[v] = struct{}{}
	}
	if len(set1) > len(set2) {
		set1, set2 = set2, set1
	}
	for v := range set1 {
		if _, has := set2[v]; has {
			intersection = append(intersection, v)
		}
	}
	return
}
