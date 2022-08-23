package chessboard

func movesToChessboard(board [][]int) int {
	n := len(board)
	rowSum := 0
	colSum := 0
	rowDiff := 0
	colDiff := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			// 行与行之间需要满足各个对应的格子全等或者全不等，这样子可保证在做列交换时性质不会改变
			if board[0][0]^board[0][j]^board[i][j]^board[i][0] != 0 {
				return -1
			}
		}
	}
	for i := 0; i < n; i++ {
		// 因为有了上面的遍历所确定的性质，所以可以只看第一行和第一列，并以(0,0)点为1的棋盘为参照
		rowSum += board[0][i]
		colSum += board[i][0]
		if board[i][0] == i%2 {
			rowDiff++
		}
		if board[0][i] == i%2 {
			colDiff++
		}
	}
	if n/2 > rowSum || rowSum > (n+1)/2 {
		return -1
	}
	if n/2 > colSum || colSum > (n+1)/2 {
		return -1
	}
	// rowDiff和colDiff需要是偶数，n - rowDiff/colDiff表示翻转棋盘
	if n%2 == 1 {
		if rowDiff%2 == 1 {
			rowDiff = n - rowDiff
		}
		if colDiff%2 == 1 {
			colDiff = n - colDiff
		}
	} else {
		if n-rowDiff < rowDiff {
			rowDiff = n - rowDiff
		}

		if n-colDiff < colDiff {
			colDiff = n - colDiff
		}
	}
	return (colDiff + rowDiff) / 2
}
