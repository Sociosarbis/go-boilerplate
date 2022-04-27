package atlantic

func dfs(heights, visited *[][]int, prev, i, j, flag int) {
	if i >= 0 && i < len(*heights) && j >= 0 && j < len((*heights)[0]) {
		if (*visited)[i][j]&flag == 0 && (*heights)[i][j] >= prev {
			(*visited)[i][j] |= flag

			dfs(heights, visited, (*heights)[i][j], i-1, j, flag)
			dfs(heights, visited, (*heights)[i][j], i+1, j, flag)
			dfs(heights, visited, (*heights)[i][j], i, j-1, flag)
			dfs(heights, visited, (*heights)[i][j], i, j+1, flag)
		}
	}
}

func pacificAtlantic(heights [][]int) [][]int {
	m := len(heights)
	n := len(heights[0])
	visited := make([][]int, m)
	for i := 0; i < m; i++ {
		visited[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		dfs(&heights, &visited, 0, i, 0, 1)
		dfs(&heights, &visited, 0, i, n-1, 2)
	}

	for i := 0; i < n; i++ {
		dfs(&heights, &visited, 0, 0, i, 1)
		dfs(&heights, &visited, 0, m-1, i, 2)
	}

	ret := [][]int{}

	for i, row := range visited {
		for j, cell := range row {
			if cell == 3 {
				ret = append(ret, []int{i, j})
			}
		}
	}
	return ret
}
