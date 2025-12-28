package negatives

func countNegatives(grid [][]int) int {
	j := len(grid[0]) - 1
	var out int
	for _, row := range grid {
		out += len(row) - j - 1
		for ; j >= 0; j-- {
			if row[j] >= 0 {
				break
			} else {
				out++
			}
		}
	}
	return out
}
