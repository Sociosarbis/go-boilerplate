package keys

var DIRS = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func shortestPathAllKeys(grid []string) int {
	target := 1
	bfs := [][]int{}
	r := len(grid)
	c := len(grid[0])
	visited := make([][]map[int]bool, r)
	for i, row := range grid {
		visited[i] = make([]map[int]bool, c)
		for j, cell := range row {
			visited[i][j] = map[int]bool{}
			if cell >= 97 && cell <= 122 {
				target |= 1 << (cell - 96)
			}
			if cell == '@' {
				bfs = append(bfs, []int{i, j, 1})
			}
		}
	}
	if target == 1 {
		return 0
	}
	ret := 0
	visited[bfs[0][0]][bfs[0][1]][1] = true
	for len(bfs) != 0 {
		ret++
		n := len(bfs)
		for i := 0; i < n; i++ {
			item := bfs[i]
		loop:
			for _, dir := range DIRS {
				x := item[0] + dir[0]
				y := item[1] + dir[1]
				if x >= 0 && x < r && y >= 0 && y < c {
					mask := item[2]
					if grid[x][y] != '#' && !visited[x][y][mask] {
						if grid[x][y] >= 97 && grid[x][y] <= 122 {
							mask |= 1 << (grid[x][y] - 96)
							if mask == target {
								return ret
							}
						} else if grid[x][y] >= 65 && grid[x][y] <= 90 {
							if mask&(1<<int(grid[x][y]-64)) == 0 {
								continue loop
							}
						}
						visited[x][y][mask] = true
						bfs = append(bfs, []int{x, y, mask})
					}
				}
			}
		}
		bfs = bfs[n:]
	}
	return -1
}
