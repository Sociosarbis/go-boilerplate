package move

func checkMove(board [][]byte, rMove int, cMove int, color byte) bool {
	pos := [2]int{}
	m := len(board)
	n := len(board[0])
	var oppositeColor byte
	if color == 'B' {
		oppositeColor = 'W'
	} else {
		oppositeColor = 'B'
	}
	for i := -1; i < 2; i++ {
	loop:
		for j := -1; j < 2; j++ {
			if i != 0 || j != 0 {
				pos[0], pos[1] = rMove+i, cMove+j
				var count int
				for pos[0] >= 0 && pos[0] < m && pos[1] >= 0 && pos[1] < n {
					if board[pos[0]][pos[1]] == oppositeColor {
						count++
					} else if board[pos[0]][pos[1]] == color {
						if count != 0 {
							return true
						} else {
							continue loop
						}
					} else {
						continue loop
					}
					pos[0] += i
					pos[1] += j
				}
			}
		}
	}
	return false
}
