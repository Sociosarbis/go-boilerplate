package area

func minimumArea(grid [][]int) int {
	l, r, t, b := len(grid[0]), 0, len(grid), 0
	for i, row := range grid {
		for j, cell := range row {
			if cell == 1 {
				if i < l {
					l = i
				}
				if i > r {
					r = i
				}
				if j < t {
					t = j
				}
				if j > b {
					b = j
				}
			}
		}
	}
	if l > r {
		return 0
	}
	return (r - l + 1) * (b - t + 1)
}
