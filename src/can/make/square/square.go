package square

func canMakeSquare(grid [][]byte) bool {
	for i := 1; i < 3; i++ {
		for j := 1; j < 3; j++ {
			var b int
			var w int
			for ii := -1; ii < 1; ii++ {
				for jj := -1; jj < 1; jj++ {
					if grid[i+ii][j+jj] == 'W' {
						w++
					} else {
						b++
					}
				}
			}
			if b >= 3 || w >= 3 {
				return true
			}
		}
	}
	return false
}
