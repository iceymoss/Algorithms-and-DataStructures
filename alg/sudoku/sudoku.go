package main

func isValidSudoku(board [][]byte) bool {
	//对每一层遍历
	mMin := make([]map[byte]string, 3)
	rows := make(map[byte]int)
	cols := make(map[byte]int)
	for row := 0; row < 9; row++ {

		//填充三个九宫格
		if row%3 == 0 {
			for i := 0; i < 3; i++ {
				mMin[i] = make(map[byte]string)
			}
		}

		for col := 0; col < 9; col++ {
			//判断行
			if board[row][col] != '.' {
				if _, ok := rows[board[row][col]]; ok && rows[board[row][col]] == row {
					return false
				} else {
					rows[board[row][col]] = row
				}
			}

			//判断列
			if board[col][row] != '.' {
				if _, ok := cols[board[col][row]]; ok && cols[board[col][row]] == row {
					return false
				} else {
					cols[board[col][row]] = row
				}
			}

			//判断九宫格
			if board[row][col] != '.' {
				if _, ok := mMin[col/3][board[row][col]]; ok {
					return false
				} else {
					mMin[col/3][board[row][col]] = ""
				}
			}
		}
	}
	return true
}

func main() {

}
