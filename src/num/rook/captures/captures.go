package captures

func numRookCaptures(board [][]byte) int {
	for i, r := range board {
		for j, c := range r {
			if c == 'R' {
				var ret int
				dirs := [4][2]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
			loop:
				for _, dir := range dirs {
					x, y := i+dir[0], j+dir[1]
					for x >= 0 && x < 8 && y >= 0 && y < 8 {
						switch board[x][y] {
						case 'p':
							ret++
							continue loop
						case 'B':
							continue loop
						}
						x += dir[0]
						y += dir[1]
					}
				}
				return ret
			}
		}
	}
	return 0
}
