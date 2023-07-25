package ArrayAndString

func transpose(matrix [][]int) [][]int {
	res := make([][]int, len(matrix[0]))
	for i := 0; i < len(res); i++ {
		res[i] = make([]int, len(matrix))
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			res[j][i] = matrix[i][j]
		}
	}
	return res
}
