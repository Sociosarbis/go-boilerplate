package area

func projectionArea(grid [][]int) int {
	xy := 0
	yz := 0
	for _, row := range grid {
		temp := 0
		for _, cell := range row {
			if cell != 0 {
				xy++
				if cell > temp {
					temp = cell
				}
			}
		}
		yz += temp
	}

	xz := 0
	for i := 0; i < len(grid); i++ {
		temp := 0
		for j := 0; j < len(grid); j++ {
			if grid[j][i] > temp {
				temp = grid[j][i]
			}
		}
		xz += temp
	}

	return xy + yz + xz
}
