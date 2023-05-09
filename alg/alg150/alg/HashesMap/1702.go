package HashesMap

func UniqueOccurrences(arr []int) bool {
	//使用m记录数组中某个数出现的次数， tagm记录m中的次数
	m, tagm := make(map[int]int), make(map[int]bool)
	for _, v := range arr {
		if _, ok := m[v]; !ok {
			m[v] = 1
		} else {
			m[v] += 1
		}
	}

	for _, v := range m {
		if !tagm[v] {
			tagm[v] = true
		} else {
			return false
		}
	}
	return true
}
