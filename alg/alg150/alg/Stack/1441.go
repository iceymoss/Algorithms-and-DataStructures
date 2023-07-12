package Stack

// 思路：从1开始到n，每次拿出一个数x，判断x是否在target中，如果在则是push, 如果不在说明被push和pop了
func buildArray(target []int, n int) []string {
	operate := make([]string, 0)
	start := 1
	for _, num := range target {
		for i := start; i != num; i++ {
			operate = append(operate, "Push", "Pop")
		}
		operate = append(operate, "Push")
		start = num + 1
	}
	return operate
}
