package rotting

func orangesRotting(grid [][]int) int {
	bfs := [][2]int{}
	var remain int
	for i, row := range grid {
		for j, cell := range row {
			switch cell {
			case 2:
				bfs = append(bfs, [2]int{i, j})
			case 1:
				remain++
			}
		}
	}

	c := len(bfs)
	m := len(grid)
	n := len(grid[0])
	var ret int
	for remain != 0 && c != 0 {
		for i := 0; i < c; i++ {
			x, y := bfs[i][0], bfs[i][1]
			options := [4][2]int{{x - 1, y}, {x + 1, y}, {x, y - 1}, {x, y + 1}}
			for _, o := range options {
				ii, jj := o[0], o[1]
				if ii >= 0 && ii < m && jj >= 0 && jj < n && grid[ii][jj] == 1 {
					grid[ii][jj] = 2
					bfs = append(bfs, [2]int{ii, jj})
					remain--
				}
			}
		}
		ret++
		bfs = bfs[c:]
		c = len(bfs)
	}
	if remain != 0 {
		return -1
	}
	return ret
}
