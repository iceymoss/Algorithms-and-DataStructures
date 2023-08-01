package HashesMap

//直接哈希表
func singleNumber(nums []int) int {
	tag := make(map[int]int)
	for _, v := range nums {
		tag[v]++
	}
	for i, v := range tag {
		if v == 1 {
			return i
		}
	}
	return 0
}
