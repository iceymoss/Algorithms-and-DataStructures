package main

func exist(board [][]byte, word string) bool {
	//判断数据有效性
	length, count := len(board), len(board[0])
	if length == 0 || count == 0 {
		return false
	}

	//使用一个二维数组记录对应位置是否被访问过
	isVisited := make([][]bool, length)
	for i := 0; i < length; i++ {
		isVisited[i] = make([]bool, count)
	}

	var find func(r, c, i int) bool
	find = func(r, c, i int) bool {
		//递归条件：如果word遍历完了，说明word存在board中
		if i == len(word) {
			return true
		}

		//判断当前查找的位置是否越界
		if r < 0 || r >= length || c < 0 || c >= count {
			return false
		}

		//判断当期位置是否被访问or是否和目标值相同
		if isVisited[r][c] || board[r][c] != word[i] {
			return false
		}

		//标记当前位置
		isVisited[r][c] = true
		//递归到下一层，查找当前位置的相邻位置是否有目标值
		isfindvalue := find(r-1, c, i+1) || find(r+1, c, i+1) || find(r, c-1, i+1) || find(r, c+1, i+1)

		//开始回溯：
		//递归返回到当前层，做后续处理
		//判断查询当前节点的相邻节点的返回结果
		//如果找到相邻节点，直接返回到上一层
		if isfindvalue {
			return true
		} else { //没有找到目标值，需要做恢复处理
			isVisited[r][c] = false
			return false
		}
	}

	//遍历二维数组
	for i := 0; i < length; i++ {
		for j := 0; j < count; j++ {
			if board[i][j] == word[0] && find(i, j, 0) {
				return true
			}
		}
	}
	return false
}
