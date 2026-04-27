package path

// [path_id][to]
var dirs [6][4]bool = [6][4]bool{
	{false, true, false, true},
	{true, false, true, false},
	{false, false, true, true},
	{false, true, true, false},
	{true, false, false, true},
	{true, true, false, false},
}

var idToDir = [4][2]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

func hasValidPath(grid [][]int) bool {
	m, n := len(grid), len(grid[0])
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]bool, n)
	}
	bfs := [][2]int{{0, 0}}
	visited[0][0] = true
	l := 1
	for l > 0 {
		for i := 0; i < l; i++ {
			r, c := bfs[i][0], bfs[i][1]
			for j, ok := range dirs[grid[r][c]-1] {
				if !ok {
					continue
				}
				dir := idToDir[j]
				nr, nc := r+dir[0], c+dir[1]
				if nr >= 0 && nr < m && nc >= 0 && nc < n && !visited[nr][nc] {
					ndir := dirs[grid[nr][nc]-1]
					if ndir[(j+2)%4] {
						bfs = append(bfs, [2]int{nr, nc})
						visited[nr][nc] = true
					}
				}
			}
		}
		bfs = bfs[l:]
		l = len(bfs)
	}
	return visited[m-1][n-1]
}
