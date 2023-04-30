package main

import "fmt"

func exist1(board [][]byte, word string) bool {
	lenght, count := len(board), len(board[0])
	if lenght == 0 || count == 0 {
		return false
	}
	if len(word) == 0 {
		return false
	}

	visited := make([][]bool, lenght)
	for i := 0; i < lenght; i++ {
		visited[i] = make([]bool, count)
	}

	//对二维数组遍历
	for i := 0; i < lenght; i++ {
		for j := 0; j < count; j++ {
			//对字符串进行遍历
			y, x := i, j
			var res bool
			for k := 0; k < len(word)-1; k++ {
				if board[y][x] == word[k] {
					visited[y][x] = true
					y, x, res = dfs(board, visited, word, k+1, y, x)
				} else {
					res = false
					break
				}
			}
			if res {
				return true
			}
		}
	}
	return false
}

// 遍历当前位置的上下左右
func dfs(board [][]byte, visited [][]bool, word string, index, i, j int) (int, int, bool) {
	ishave := true
	if i-1 >= 0 && board[i-1][j] == word[index] && !visited[i-1][j] {
		i, j = i-1, j
	} else if i+1 < len(board) && board[i+1][j] == word[index] && !visited[i+1][j] {
		i, j = i+1, j
	} else if j-1 >= 0 && board[i][j-1] == word[index] && !visited[i][j-1] {
		i, j = i, j-1
	} else if j+1 < len(board[0]) && board[i][j+1] == word[index] && !visited[i][j+1] {
		i, j = i, j+1
	} else {
		ishave = false
	}
	return i, j, ishave
}

func main() {
	b := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	word := "ABCB"
	fmt.Println(exist(b, word))
}
