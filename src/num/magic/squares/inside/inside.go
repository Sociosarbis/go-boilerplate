package inside

func isMagicSquare(grid [][]int, left, top int) bool {
	var mask int
	var sum int
	for i := 0; i < 3; i++ {
		sum += grid[top][left+i]
	}
	for i := 0; i < 3; i++ {
		row := grid[top+i]
		var temp int
		for j := 0; j < 3; j++ {
			value := row[left+j]
			if value < 1 || value > 9 {
				return false
			}
			if mask&(1<<value) != 0 {
				return false
			}
			mask |= 1 << value
			temp += value
		}
		if temp != sum {
			return false
		}
	}
	for i := 0; i < 3; i++ {
		var temp int
		for j := 0; j < 3; j++ {
			temp += grid[top+j][left+i]
		}
		if temp != sum {
			return false
		}
	}
	var temp int
	for i := 0; i < 3; i++ {
		temp += grid[top+i][left+i]
	}
	if temp != sum {
		return false
	}
	temp = 0
	for i := 0; i < 3; i++ {
		temp += grid[top+i][left+2-i]
	}
	return temp == sum
}

func numMagicSquaresInside(grid [][]int) int {
	m, n := len(grid)-2, len(grid[0])-2
	var out int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if isMagicSquare(grid, j, i) {
				out++
			}
		}
	}
	return out
}
