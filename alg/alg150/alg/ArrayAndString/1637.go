package ArrayAndString

import "sort"

func maxWidthOfVerticalArea(points [][]int) int {
	//按照x进行排序，然后计算相邻点直接的x差值即可
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})
	sumMax := 0
	tag := points[0]
	for i, v := range points {
		sumMax = max(sumMax, points[i][0]-tag[0])
		tag = v
	}
	return sumMax
}

//func max(x, y int) int {
//	if x > y {
//		return x
//	}
//	return y
//}
