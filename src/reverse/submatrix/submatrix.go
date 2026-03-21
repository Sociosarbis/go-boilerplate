package submatrix

func reverseSubmatrix(grid [][]int, x int, y int, k int) [][]int {
	temp := make([]int, k)
	end := k / 2
	for i := 0; i < end; i++ {
		copy(temp, grid[x+i][y:y+k])
		copy(grid[x+i][y:y+k], grid[x+k-1-i][y:y+k])
		copy(grid[x+k-1-i][y:y+k], temp)
	}
	return grid
}
