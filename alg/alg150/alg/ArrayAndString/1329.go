package ArrayAndString

import "sort"

func diagonalSort(mat [][]int) [][]int {
	//思路：使用map记录每一个对角线的数组，d[i-j]中i-j就表示对应对角线的key,因为对角线上的值i-j总相等，然后对数组进行排序
	m, n := len(mat), len(mat[0])
	d := make(map[int][]int, m+n-1)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			d[i-j] = append(d[i-j], mat[i][j])
		}
	}
	for _, v := range d {
		sort.Ints(v)
	}

	//it用来记录d[i-j]中数组的的索引
	it := make(map[int]int, m+n-1)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			mat[i][j] = d[i-j][it[i-j]]
			it[i-j]++
		}
	}
	return mat
}
