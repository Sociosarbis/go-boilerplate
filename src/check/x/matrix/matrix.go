package matrix

func checkXMatrix(grid [][]int) bool {
	last := len(grid) - 1
	for i, row := range grid {
		for j, v := range row {
			if i == j || i == last-j {
				if v == 0 {
					return false
				}
			} else {
				if v != 0 {
					return false
				}
			}
		}
	}
	return true
}
